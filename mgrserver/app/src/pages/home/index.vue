
<template>
  <div>
    <van-nav-bar title="宝宝积分">
      <template #right>
        <van-icon name="ellipsis" size="18" />
      </template>
    </van-nav-bar>

    <van-cell-group>
      <van-cell
        is-link
        clickable
        center
        url="/score/add"
        v-for="item in dataList.items"
        v-bind:key="item.uid"
      >
        <!-- 使用 title 插槽来自定义标题 -->
        <template>
          <van-row type="flex" align="center">
            <van-col span="4" align="center">
              <van-image
                :src="item.url"
                round
                fit="cover"
                width="4rem"
                height="4rem"
              >
              </van-image>
            </van-col>
            <van-col span="3"> </van-col>
            <van-col span="12">
              <div class="h5">{{ item.name }}</div>
              <div class="text-muted">{{ getGrowAge(item.birthday) }}</div>
              <div>
                总分:<span class="text-success h4"> {{ item.score }}</span> 分
              </div>
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
    getGrowAge(birthday) {
      var now = new Date();
      var year = now.getFullYear();
      var month = now.getMonth() + 1;
      var day = now.getDate();
      var hour = now.getHours();
      var minute = now.getMinutes();
      var second = now.getSeconds();

      var myDate = new Date(birthday);
      var myYear = myDate.getFullYear();
      var myMonth = myDate.getMonth() + 1;
      var myDay = myDate.getDate();
      var myHour = myDate.getHours();
      var myMinute = myDate.getMinutes();
      var mySecond = myDate.getSeconds();

      var gapSecond = second - mySecond;
      if (gapSecond < 0) {
        minute -= 1;
        gapSecond = 60 - mySecond + second;
      }
      var gapMinute = minute - myMinute;
      if (gapMinute < 0) {
        hour -= 1;
        gapMinute = 60 - myMinute + minute;
      }
      var gapHour = hour - myHour;
      if (gapHour < 0) {
        day -= 1;
        gapHour = 24 - myHour + hour;
      }
      var gapDay = day - myDay;
      if (gapDay < 0) {
        month -= 1;
        gapDay = this.getDaysOfMonth(birthday) - myDay + day;
      }
      var gapMonth = month - myMonth;
      if (gapMonth < 0) {
        year -= 1;
        gapMonth = 12 - myMonth + month;
      }
      var gapYear = year - myYear;
      if (gapYear < 0) {
        gapYear = 0;
      }
      var dateStr = gapYear + "岁 " + gapMonth + "个月" + myDay + "天";
      return dateStr;
    },
    getDaysOfMonth(dateStr) {
      var date = new Date(dateStr);
      var year = date.getFullYear();
      var mouth = date.getMonth() + 1;
      var day = 0;

      if (mouth == 2) {
        day = this.isLeapYear(year) ? 29 : 28;
      } else if (
        mouth == 1 ||
        mouth == 3 ||
        mouth == 5 ||
        mouth == 7 ||
        mouth == 8 ||
        mouth == 10 ||
        mouth == 12
      ) {
        day = 31;
      } else {
        day = 30;
      }
      return day;
    },
    isLeapYear(year) {
      return (year % 4 == 0 && year % 100 != 0) || year % 400 == 0;
    },
  },
};
</script>
