<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="字典配置" name="DictInfoDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.id | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">名称:</div>
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
                    <div class="pull-right" style="margin-right: 10px">值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.value | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">类型:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.type | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("status") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">排序值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.sort_no |  fltrNumberFormat(0)}}</div>
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
        tabName: "DictInfoDetail",
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
        this.info = this.$http.xget("/dict/info",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>