
<template>
<div>
<router-view/>

<van-tabbar route >
   <van-tabbar-item replace to="/home" icon="home-o">首页</van-tabbar-item>
  <van-tabbar-item replace to="/search" icon="search">积分记录</van-tabbar-item>
  <van-tabbar-item replace to="/search" icon="search">个人中心</van-tabbar-item>
</van-tabbar>
</div>
</template>
<script>
export default {
  data() {
    return {
       active: 'home',
    }
  },
  created() {

  },
  mounted() {

  },
  methods: {
  }
}
</script>
