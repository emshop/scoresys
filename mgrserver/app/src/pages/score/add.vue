
<template>
  <div>
    <van-skeleton title avatar :row="3" :loading="loading">
    <van-nav-bar title="记录一笔" left-arrow @click-left="onClickLeft" />
    <van-cell-group title="奖罚类型">
      <van-cell>
        <van-tree-select
          :items="items"
          :active-id.sync="activeIds"
          :main-active-index.sync="activeIndex"
          @click-item="onSelectChanged"
        />
      </van-cell>
    </van-cell-group>

    <van-cell-group title="分数信息">
      <van-form>
        <van-field
          readonly
          clickable
          name="姓名"
          label-align="left"
          input-align="left"
          :value="user.text"
          label="姓名"
          placeholder="点击选择用户"
          @click="showPicker = true"
        />
        <van-popup v-model="showPicker" position="bottom">
          <van-picker show-toolbar @change="onChange" :columns="users"   @cancel="showPicker=false" @confirm="showPicker=false" />
        </van-popup>

        <van-field
          v-model="reward_score"
          type="number"
          input-align="left"
          label-align="left"
          name="奖分"
          label="奖分"
          readonly
          placeholder="0"
        />
        <van-field
          v-model="pub_score"
          type="number"
          input-align="left"
          label-align="left"
          readonly
          name="罚分"
          label="罚分"
          placeholder="0"
        />

        <van-submit-bar
          :price="total"
          label="总分："
          ext-align="left"
          button-text="提交分数"
          @submit="onSubmit"
        />
      </van-form>
    </van-cell-group>
    </van-skeleton>
  </div>
</template>
<script>
export default {
  data() {
    return {
      loading:true,
      reward_score: 0,
      pub_score: 0,
      showPicker: false,
      total: 0,
      user: {},
      users: [],
      rewardItem: [],
      punItem: [],
      paging: { ps: 10, pi: 1 },
      items: [
        { text: "奖分", children: [{ text: "7点前起床" }] },
        { text: "罚分", children: [] },
      ],
      activeIds: [],
      activeIndex: 0,
    };
  },
  created() {},
  mounted() {
    this.query();
    this.getUsers();
  },
  methods: {
    getUsers() {
      let that = this     
      this.$http.post("/user/info/query", this.paging).then(res=>{
        let users = [];
        res.items.forEach(function (item) {
          users.push({ text: item.name, id: item.uid });
        });
        that.users = users;
        that.user = users[0];
        if (that.$route.params.id){
          that.users.forEach(function(u){
            if (u.id==that.$route.params.id){
                that.user = u
            }
          })
        }
        that.loading=false
      })
     
    },
    onClickLeft() {
      this.$router.push("/home/index");
    },
    query() {
      let rewardItem = this.rewardItem;
      let reward = this.$enum.get("reward_info") || [];
      reward.forEach(function (item) {
        rewardItem.push({ text:item.name+"("+  item.value +"分)", id: item.id, value: item.value });
      });

      let punItem = this.punItem;
      let pun = this.$enum.get("punishment_info") || [];
      pun.forEach(function (item) {
        punItem.push({ text: item.name, id: item.id, value: item.value });
      });
      this.rewardItem = rewardItem;
      this.punItem = punItem;
      this.items[0].children = rewardItem;
      this.items[1].children = punItem;
    },
    getScore(items, id) {
      var score = 0;
      items.forEach(function (item) {
        if (parseInt(item.id) == id) {
          score = parseInt(item.value);
          return false;
        }
      });
      return score;
    },
    onChange(p, value, i) {
      this.showPicker = false;
      this.user = value;

    },
    onSubmit() {
      let input={uid:this.user.id,score:this.total/100}
      let that=this
      this.$http.post("/score/record/add", input).then(v=>{
        that.$msg.success("保存成功")
        that.$router.push("/home/index");
      }).catch(ex=>{
         that.$msg.fail("保存失败")
      })


    },
    onSelectChanged(data) {
      var reward = 0;
      var pum = 0;
      var that = this;
      this.activeIds.forEach(function (item) {
        let score = that.getScore(that.rewardItem, item);
        if (score > 0) {
          reward += score;
          return;
        }
        score = that.getScore(that.punItem, item);
        pum += score;
      });
      this.reward_score = reward;
      this.pub_score = pum;
      this.total = (this.reward_score + this.pub_score) * 100;
    },
  },
};
</script>
