
<template>
<div>
<router-view/>
<van-tabbar route >
   <van-tabbar-item replace to="/home/index" icon="home-o">首页</van-tabbar-item>
  <van-tabbar-item replace to="/home/score" icon="certificate">记录</van-tabbar-item>
  <van-tabbar-item replace to="/home/exchange" icon="point-gift-o">兑换</van-tabbar-item>
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
