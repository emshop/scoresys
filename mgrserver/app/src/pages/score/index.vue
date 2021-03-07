<template>
  <div>
    <van-nav-bar :title="title" @click-left="onClickLeft" left-arrow />
    <van-tabs v-model="active" @change="onChange" animated swipeable>
      <van-tab title="奖分">
        <van-cell-group title="选择项目">
          <van-cell>
            <van-button
              plain
              :type="item.type"
              style="margin-right: 10px"
              class="badge-icon"
              size="small"
              :class="item.selected"
              @click="onClick(item)"
              v-for="item in rewardItem"
              v-bind:key="item.id"
              >{{ item.text }}</van-button
            >
          </van-cell>
        </van-cell-group>
      </van-tab>
      <van-tab title="罚分">
        <van-cell-group title="选择项目">
          <van-cell>
            <van-button
              plain
              :type="item.type"
              style="margin-right: 10px"
              class="badge-icon"
              size="small"
              :class="item.selected"
              @click="onClick(item)"
              v-for="item in punItem"
              v-bind:key="item.id"
              >{{ item.text }}</van-button
            >
          </van-cell>
        </van-cell-group>
      </van-tab>
    </van-tabs>
    <van-submit-bar
      label="总分："
      :price="pub_score * 100"
      ext-align="left"
      button-text="提交"
      @submit="onSubmit"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: "",
      user: {},
      active: 0,
      pub_score: 0,
      rewardItem: [],
      selected: {},
      punItem: [],
      paging: { ps: 10, pi: 1 },
    };
  },
  created() {},
  mounted() {
    this.getUsers();
    this.queryReward();
    this.queryPub();

  },
  methods: {
    onSubmit() {
      let that = this;
      let ids = [];
      let input = { uid: that.$route.params.id};
      for (var item in this.selected) {
        ids.push(item);
      }
      input.ids = ids.toString();
      this.$dialog
        .confirm({
          title: "提交",
          message: "确认提交分数信息吗?",
        })
        .then(() => {
          this.$http
            .post("/score/record/add", input)
            .then((v) => {
              that.$msg.success("保存成功");
              that.$router.push("/home/index");
            })
            .catch((ex) => {
              that.$msg.fail("保存失败");
            });
        })
        .catch((e) => {});
    },
    onClickLeft() {
      this.$router.push("/home/index");
    },
    onChange(i) {
      this.active = i;
      this.title =
        "给 " + this.user.name + (this.active == 0 ? " 奖分" : " 罚分");
    },
    getUsers() {
      let that = this;
      let input = { uid: that.$route.params.id };
      this.$http.get("/user/info", input).then((res) => {
        that.user = res;
        that.title = "给 " + res.name + (that.active == 0 ? " 奖分" : " 罚分");
      });
    },

    onClick(item) {
      if (this.selected[item.id]) {
        delete this.selected[item.id];
        item.type = "default";
        this.pub_score = parseInt(this.pub_score) - parseInt(item.value);
      } else {
        this.selected[item.id] = item.id;
        item.type = "danger";
        this.pub_score = parseInt(this.pub_score) + parseInt(item.value);
      }
    },
    queryReward() {
      let rewardItem = this.rewardItem;
      let reward = this.$enum.get("reward_info") || [];
      reward.forEach(function (item) {
        rewardItem.push({
          text: item.name + "(" + item.value + "分)",
          id: item.id,
          type: "default",
          value: item.value,
        });
      });
      this.rewardItem = rewardItem;
    },
    queryPub() {
      let punItem = this.punItem;
      let pun = this.$enum.get("punishment_info") || [];
      pun.forEach(function (item) {
        if (item.value > 0) {
          item.value = -1 * item.value;
        }
        punItem.push({
          text: item.name,
          id: item.id,
          value: item.value,
          type: "default",
        });
      });

      this.punItem = punItem;
    },
  },
};
</script>

<style>
</style>