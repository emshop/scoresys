
<template>
  <div>
    <van-nav-bar
      title="积分信息"
    />
    <van-cell-group>
      <van-cell
        is-link
        size="large"
        v-for="item in dataList.items"
        v-bind:key="item.uid"
      >
        <!-- 使用 title 插槽来自定义标题 -->
        <template>
          <van-row type="flex" align="top" justify="center">
            <van-col span="8" align="center">
              <van-image :src="item.url" width="80">               
              </van-image>
            </van-col>
            <van-col span="1"> </van-col>
            <van-col span="12">
              <div>{{ item.name }}</div>
              <div> {{ item.score }}分</div>
            </van-col>
          </van-row>
        </template>
      </van-cell>
    </van-cell-group>
  </div>
</template>
<script>
export default {
  data() {
    return {
      paging: { ps: 10, pi: 1 },
      dataList: { items: [] },
    };
  },
  created() {},
  mounted() {
    this.query();
  },
  methods: {
    query() {
      let res = this.$http.xpost("/user/info/query", this.paging);
      this.dataList.items = res.items;
      this.dataList.count = res.count;
    },
  },
};
</script>
