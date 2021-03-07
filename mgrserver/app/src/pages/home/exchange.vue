<template>
  <div>
    <van-nav-bar
      title="积分兑换"
      class="bg-danger"
      @click-left="onClickLeft"
      left-arrow
    >
      <template #right>
        <van-icon name="ellipsis" size="18" />
      </template>
    </van-nav-bar>
    <van-form @submit="onSubmit">
      <van-cell-group title="请填写兑换信息">
        <van-field
          readonly
          clickable
          name="宝宝"
          label-align="left"
          input-align="right"
          :value="user.text"
          label="宝宝"
          placeholder="点击选择用户"
          @click="showPicker = true"
        />
        <van-popup v-model="showPicker" position="bottom">
          <van-picker
            show-toolbar
            @change="onChange"
            :columns="users"
            @cancel="showPicker = false"
            @confirm="showPicker = false"
          />
        </van-popup>

        <van-field
          rows="5"
          autosize
          type="textarea"
          maxlength="20"
          show-word-limit
          v-model="input.content"
          input-align="right"
          label-align="left"
          name="内容"
          label="内容"
          :rules="[{ required: true, message: '请输入兑换内容' }]"
        />
        <van-field
          v-model="input.score"
          type="number"
          input-align="right"
          label-align="left"
          style="font-size: 16px"
          name="分数"
          label="分数"
          :rules="[{ required: true, message: '请填写兑换分数' }]"
        />
      </van-cell-group>

      <div style="margin: 16px">
        <van-button round block type="info" native-type="submit"
          >提交</van-button
        >
      </div>
    </van-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      input: {},
      paging: { ps: 10, pi: 1 },
      showPicker: false,
      user: {},

      users: [],
    };
  },
  created() {},
  mounted() {
    this.getUsers();
  },
  methods: {
    onSubmit() {
      this.input.uid = this.user.id;
      let that = this;
      this.$dialog
        .confirm({
          title: "提交",
          message: "确认提交兑换信息吗?",
        })
        .then(() => {
          this.$http
            .post("/score/record/exchange", that.input)
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
    onChange(p, value, i) {
      this.showPicker = false;
      this.user = value;
    },
    onClickLeft() {
      this.$router.push("/home/index");
    },
    getUsers() {
      let that = this;
      this.$http.post("/user/info/query", this.paging).then((res) => {
        let users = [];
        res.items.forEach(function (item) {
          users.push({ text: item.name, id: item.uid });
        });
        that.users = users;
        that.user = users[0];
      });
    },
  },
};
</script>
<style lang="css" scoped>
/deep/ .van-nav-bar .van-icon {
  color: #fff;
}

/deep/ .van-nav-bar__title {
  color: #fff;
}
</style>