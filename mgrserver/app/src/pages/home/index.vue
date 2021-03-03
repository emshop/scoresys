
<template>
  <div>
    <van-nav-bar
      title="积分信息"
    >
      <template #right>
        <van-icon name="ellipsis" size="18"/>
  </template>
    </van-nav-bar>
    <van-cell-group title="宝宝信息" >
      <van-cell
        is-link
        clickable
        center
         url="/score" 
        v-for="item in dataList.items"
        v-bind:key="item.uid"
      >
        <!-- 使用 title 插槽来自定义标题 -->
        <template>
          <van-row type="flex" align="top">
            <van-col span="4" align="center">
              <van-image :src="item.url" round   fit="cover" width="4rem" height="4rem">               
              </van-image>
            </van-col>
            <van-col span="3"> </van-col>
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
