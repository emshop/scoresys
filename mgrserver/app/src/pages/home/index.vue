
<template>
  <div>
    <van-nav-bar title="宝宝积分" class="bg-danger">
      <template #right>
        <van-icon name="ellipsis" size="18" />
      </template>
    </van-nav-bar>

    <van-cell-group title="我的宝宝">
      <van-cell
        is-link
        clickable
        center
        :url="'/score/add/' + item.uid"
        v-for="item in dataList.items"
        v-bind:key="item.uid"
      >
        <!-- 使用 title 插槽来自定义标题 -->
        <template>
          <van-row type="flex" align="center">
            <van-col span="5" align="center">
              <van-image
                :src="item.url"
                round
                fit="cover"
                width="4.5rem"
                height="4.5rem"
              >
              </van-image>
            </van-col>
            <van-col span="3"> </van-col>
            <van-col span="12">
              <div class="h6">{{ item.name }}</div>
              <div class="text-muted">{{ getGrowAge(item.birthday) }}</div>
              <div>
                总分:<span class="text-danger h6"> {{ item.score }}</span> 分
              </div>
            </van-col>
          </van-row>
        </template>
      </van-cell>
    </van-cell-group>

    <van-cell-group title="最近记录">
      <van-collapse
        v-for="r in records"
        v-bind:key="r.id"
        v-model="r.id"
        accordion
      >
        <van-collapse-item :title="r.group" v-name="r.id" toggle="false">
          <van-cell
            center
            v-for="rcd in r.items"
            v-bind:key="rcd.id"
            :value-class="rcd.type == 1 ? 'text-danger' : 'text-success'"
            :title="rcd.score > 0 ? '+' + rcd.score + '分' : rcd.score + '分'"
            :value="rcd.type == 1 ? '奖励' : '兑换'"
            :label="rcd.label"
          />
        </van-collapse-item>
      </van-collapse>
    </van-cell-group>
  </div>
</template>
<script>
export default {
  data() {
    return {
      activeNames: "1",
      paging: { ps: 10, pi: 1 },
      dataList: { items: [] },
      records: [
        {
          group: "2020-03-06",
          id: "1",
          items: [
            { id: 1, score: 5, type: 1, label: "超级乖宝宝" },
            { id: 1, score: -10, type: 3, label: "游戏币" },
          ],
        },
      ],
    };
  },
  created() {},
  mounted() {
    this.query();
  },
  methods: {
    query() {
      this.$http.post("/user/info/query", this.paging).then((res) => {
        this.dataList.items = res.items;
        this.dataList.count = res.count;
      });
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

<style lang="css" scoped>
/deep/ .van-ellipsis {
  color: #fff;
}

/deep/ .van-nav-bar .van-icon {
  color: #fff;
}

/deep/ .van-nav-bar__title {
  color: #fff;
}
</style>