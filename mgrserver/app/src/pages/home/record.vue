<template>
  <div>
    <van-nav-bar title="积分记录" class="bg-danger">
      <template #right>
        <van-icon name="ellipsis" size="18" />
      </template>
    </van-nav-bar>
    <van-tabs @click="onClick">
      <van-tab
        v-for="u in users"
        v-bind:key="u.uid"
        :name="u.uid"
        :title="u.name"
      ></van-tab>
    </van-tabs>

<van-pull-refresh v-model="isLoading" @refresh="onRefresh">

    <van-cell-group :title="r.group" v-for="r in records" v-bind:key="r.id">
      <van-cell
        center
        v-for="rcd in r.items"
        v-bind:key="rcd.id"
        :value-class="rcd.c_tp == '0' ? 'text-danger' : 'text-success'"
        :title="rcd.score > 0 ? '+' + rcd.score + '分' : rcd.score + '分'"
        :value="rcd.c_tp == '0' ? '奖励' : '兑换'"
        :label="rcd.content"
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
      selectUID: 0,
      users: [],
      paging: { ps: 10, pi: 1 },
      records: [
        {
          group: "",
          id: "1",
          items: [],
        },
      ],
    };
  },

  created() {},
  mounted() {
    this.getUsers();
  },
  methods: {
     onRefresh() {
       this.queryRecords()
     },
    onClick(name, title) {
      this.selectUID = name;
      this.queryRecords();
    },
    getUsers() {
      let that = this;
      this.$http.post("/user/info/query", this.paging).then((res) => {
        that.users = res.items;
         this.selectUID = this.users[0].uid;
        this.queryRecords();
      });
     
    },
    queryRecords() {
      let that = this;
      let recordList = {};
      var curr_date = new Date();
      curr_date.setDate(curr_date.getDate() - 10);
      this.$http
        .post("/score/record/q", {
          create_time: this.dateFormat(curr_date, "yyyy-MM-dd"),
          pi: this.paging.pi,
          ps: this.paging.ps,
          uid: this.selectUID,
        })
        .then((res) => {
          that.isLoading = false
          recordList.items = res.items;
          let lst = {};
          recordList.items.forEach((item) => {
            let dt = that.dateFormat(item.create_time, "yyyy-MM-dd");
            if (!lst[dt]) {
              lst[dt] = [];
            }
            lst[dt].push(item);
          });
          that.records=[]
          for (var dt in lst) {
            that.records.push({
              group: dt,
              id: lst[dt].id,
              items: lst[dt],
            });
          }
        });
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