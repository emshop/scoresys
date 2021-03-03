<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="用户信息" name="UserInfoDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">用户编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.uid | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">姓名:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.name && info.name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.name}}</div>
                      <div >{{ info.name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.name | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">头像:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.url && info.url.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.url}}</div>
                      <div >{{ info.url | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.url | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">出生日期:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.birthday | fltrDate("yyyy-MM-dd") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">分数:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.score |  fltrNumberFormat(0)}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.status | fltrEnum("status") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.create_time | fltrDate("yyyy-MM-dd") }}</div>
                  </el-col>
                </td>
              </tr>            
            </tbody>
          </table>
        </div>
	  </el-tab-pane>
	  </el-tabs>
	</div>
</template>

<script>
	export default {
    data(){
      return {
        tabName: "UserInfoDetail",
        info: {},
      }
    },
    mounted() {
      this.init();
    },
    created(){
    },
    methods: {
      init(){
        this.queryData()
      },
      queryData() {
        this.info = this.$http.xget("/user/info",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>