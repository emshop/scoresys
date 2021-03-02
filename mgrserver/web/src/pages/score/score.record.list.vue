<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.uid" clearable filterable class="input-cos" placeholder="请选择用户">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in uid" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.c_tp" @change="setScore(queryData.c_tp)" clearable filterable class="input-cos" placeholder="请选择类型">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in cTp" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button  type="primary" @click="query" size="medium">查询</el-button>
				</el-form-item>
				
				<el-form-item>
					<el-button type="success" size="medium" @click="showAdd">添加</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :max-height="maxHeight">
				
				<el-table-column   prop="c_tp" label="类型" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.c_tp | fltrEnum("c_tp")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="score" label="变动分数" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.score | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("yyyy-MM-dd") }}</div>
				</template>
				</el-table-column>
				<el-table-column  label="操作" align="center">
					<template slot-scope="scope">
						<el-button type="text" size="mini" @click="showEdit(scope.row)">编辑</el-button>
						<el-button type="text" size="mini" @click="showDetail(scope.row)">详情</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-scrollbar>
		<!-- list end-->

		<!-- Add Form -->
		<Add ref="Add" :refresh="query"></Add>
		<!--Add Form -->

		<!-- edit Form start-->
		<Edit ref="Edit" :refresh="query"></Edit>
		<!-- edit Form end-->

		<!-- pagination start -->
		<div class="page-pagination">
		<el-pagination
			@size-change="pageSizeChange"
			@current-change="pageIndexChange"
			:current-page="paging.pi"
			:page-size="paging.ps"
			:page-sizes="paging.sizes"
			layout="total, sizes, prev, pager, next, jumper"
			:total="dataList.count">
		</el-pagination>
		</div>
		<!-- pagination end -->

	</div>
</template>


<script>
import Add from "./score.record.add"
import Edit from "./score.record.edit"
export default {
  components: {
		Add,
		Edit
  },
  data () {
		return {
			paging: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
			editData:{},                //编辑数据对象
			addData:{},                 //添加数据对象 
      queryData:{},               //查询数据对象
			uid: this.$enum.get("user_info"),
			cTp: this.$enum.get("c_tp"),
			dataList: {count: 0,items: []}, //表单数据对象,
			maxHeight: document.body.clientHeight
		}
  },
  created(){
  },
  mounted(){
		this.maxHeight = this.$utility.getTableHeight("panel-body")
    this.init()
  },
	methods:{
    /**初始化操作**/
    init(){
      this.query()
		},
    /**查询数据并赋值*/
    query(){
      this.queryData.pi = this.paging.pi
			this.queryData.ps = this.paging.ps
      let res = this.$http.xpost("/score/record/query",this.$utility.delEmptyProperty(this.queryData))
			this.dataList.items = res.items
			this.dataList.count = res.count
    },
    /**改变页容量*/
		pageSizeChange(val) {
      this.paging.ps = val
      this.query()
    },
    /**改变当前页码*/
    pageIndexChange(val) {
      this.paging.pi = val
      this.query()
    },
    /**重置添加表单*/
    resetForm(formName) {
      this.dialogAddVisible = false
      this.$refs[formName].resetFields();
		},
		showDetail(val){
			var data = {
        rc_id: val.rc_id,
      }
      this.$emit("addTab","详情"+val.rc_id,"/score/record/detail",data);
		},
    showAdd(){
      this.$refs.Add.show();
		},
    showEdit(val) {
      this.$refs.Edit.editData = val
      this.$refs.Edit.show();
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
