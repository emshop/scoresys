
<template>
<div>
<router-view/>

<van-tabbar route >
   <van-tabbar-item replace to="/home" icon="home-o">首页</van-tabbar-item>
  <van-tabbar-item replace to="/search" icon="certificate">立即积分</van-tabbar-item>
  <van-tabbar-item replace to="/search" icon="contact">个人中心</van-tabbar-item>
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
