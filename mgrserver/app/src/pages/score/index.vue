
<template>
  <div>
    <van-nav-bar title="记录一笔" left-arrow @click-left="onClickLeft" />
    <van-cell-group title="请选择加减分项">
      <van-tree-select
        :items="items"
        :active-id.sync="activeIds"
        :main-active-index.sync="activeIndex"
        @click-item="onSelectChanged"
      />
    </van-cell-group>

    <van-cell-group title="宝宝信息">
      <van-form>
        <van-field
          readonly
          clickable
          name="姓名"
          :value="user"
          label="姓名"
          placeholder="点击选择用户"
          @click="showPicker = true"
        />
        <van-popup v-model="showPicker" position="bottom">
          <van-picker show-toolbar @change="onChange" :columns="users" />
        </van-popup>

        <van-field
          v-model="reward_score"
          name="奖励"
          label="奖励"
          readonly
          placeholder="0"
        />
        <van-field
          v-model="pub_score"
          readonly
          name="惩罚"
          label="惩罚"
          placeholder="0"
        />

        <div style="margin: 16px">
          <van-submit-bar
            :price="total"
            label="总分："
            ext-align="left"
            button-text="提交分数"
            @submit="onSubmit"
          />
        </div>
      </van-form>
    </van-cell-group>
  </div>
</template>
<script>
export default {
  data() {
    return {
      reward_score: 0,
      pub_score: 0,
      showPicker: false,
      total: 0,
      user: "",
      users: [],
      rewardItem: [],
      punItem: [],
      paging: { ps: 10, pi: 1 },
      items: [
        { text: "奖励", children: [{ text: "7点前起床" }] },
        { text: "惩罚", children: [] },
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
      let res = this.$http.xpost("/user/info/query", this.paging);
      let users = [];
      res.items.forEach(function (item) {
        users.push({ text: item.name, id: item.uid });
      });
      this.users = users;
    },
    onClickLeft() {
      this.$router.push("/home/index");
    },
    query() {
      let rewardItem = this.rewardItem;
      let reward = this.$enum.get("reward_info") || [];
      reward.forEach(function (item) {
        rewardItem.push({ text: item.name, id: item.id, value: item.value });
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
      this.user = value.text;
    },
    onSubmit() {},
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
