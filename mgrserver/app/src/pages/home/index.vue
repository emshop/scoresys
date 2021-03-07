
<template>
  <div>
    <van-nav-bar title="宝宝积分" class="bg-danger">
      <template #right>
        <van-icon name="ellipsis" size="18" />
      </template>
    </van-nav-bar>
    <van-pull-refresh v-model="isLoading" @refresh="onRefresh">
      <van-cell-group title="我的宝宝">
        <van-cell
          is-link
          clickable
          center
          :url="'/score/index/' + item.uid"
          v-for="item in users.items"
          v-bind:key="item.uid"
        >
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

      <van-cell-group v-for="r in records" :title="r.group" v-bind:key="r.id">
        <van-cell
          center
          v-for="rcd in r.items"
          v-bind:key="rcd.id"
          :value-class="rcd.c_tp == '0' ? 'text-danger' : 'text-success'"
          :title="rcd.score > 0 ? '+' + rcd.score + '分' : rcd.score + '分'"
          :value="rcd.c_tp == '0' ? '奖励' : '兑换'"
          :label="rcd.content + '(' + getUserName(rcd.uid) + ')'"
        />
      </van-cell-group>
    </van-pull-refresh>
  </div>
</template>
<script>
export default {
  data() {
    return {
      isLoading: false,
      activeNames: "1",
      paging: { ps: 10, pi: 1 },
      users: { items: [] },
      records: [],
    };
  },
  created() {},
  mounted() {
    this.query();
    this.queryRecords();
  },
  methods: {
    onRefresh() {
      this.paging.pi = 1;
      this.query();
    },
    queryRecords() {
      let that = this;
      let recordList = {};
      var curr_date = new Date();
      curr_date.setDate(curr_date.getDate() - 1);
      this.$http
        .post("/score/record/q", {
          create_time: this.dateFormat(curr_date, "yyyy-MM-dd"),
          pi: this.paging.pi,
          ps: this.paging.ps,
        })
        .then((res) => {
          that.isLoading = false;
          recordList.items = res.items;
          let lst = {};
          recordList.items.forEach((item) => {
            let dt = that.dateFormat(item.create_time, "yyyy-MM-dd");
            if (!lst[dt]) {
              lst[dt] = [];
            }
            lst[dt].push(item);
          });
          that.records = [];
          for (var dt in lst) {
            that.records.push({
              group: dt,
              id: lst[dt].uid,
              items: lst[dt],
            });
          }
        });
    },

    query() {
      let that = this;
      this.$http.post("/user/info/query", this.paging).then((res) => {
        that.isLoading = false;
        that.users.items = res.items;
        that.users.count = res.count;
        this.queryRecords();
      });
    },
    getUserName(uid) {
      let name = "";
      this.users.items.forEach(function (item) {
        if (item.uid == uid) {
          name = item.name;
          return;
        }
      });
      return name;
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

    dateFormat(date, fmt) {
      if (!date) {
        return "";
      }
      let d = new Date(Date.parse(date));
      var o = {
        "M+": d.getMonth() + 1, // 月份
        "d+": d.getDate(), // 日
        "h+": d.getHours(), // 小时
        "m+": d.getMinutes(), // 分
        "s+": d.getSeconds(), // 秒
        "q+": Math.floor((d.getMonth() + 3) / 3), // 季度
        S: d.getMilliseconds(), // 毫秒
      };
      if (/(y+)/.test(fmt)) {
        fmt = fmt.replace(
          RegExp.$1,
          (d.getFullYear() + "").substr(4 - RegExp.$1.length)
        );
      }
      for (var k in o) {
        if (new RegExp("(" + k + ")").test(fmt)) {
          fmt = fmt.replace(
            RegExp.$1,
            RegExp.$1.length === 1
              ? o[k]
              : ("00" + o[k]).substr(("" + o[k]).length)
          );
        }
      }
      return fmt;
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