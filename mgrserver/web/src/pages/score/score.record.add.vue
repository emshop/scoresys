<template>
  <!-- Add Form -->
  <el-dialog title="添加分数记录" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      
			<el-form-item label="用户:" prop="uid">
				<el-select size="medium" style="width: 100%;"	v-model="addData.uid"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in uid" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="类型:" prop="c_tp">
				<el-select size="medium" style="width: 100%;"	v-model="addData.c_tp"	clearable filterable class="input-cos" placeholder="---请选择---" @change="setScore(addData.c_tp)">
					<el-option v-for="(item, index) in cTp" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="变动分数:" prop="score">
				<el-select size="medium" style="width: 100%;"	v-model="addData.score"	clearable filterable class="input-cos" placeholder="---请选择---" @change="handleChooseTool()">
					<el-option v-for="(item, index) in score" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="medium" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="medium" type="success" @click="add('addForm')">确 定</el-button>
    </div>
  </el-dialog>
  <!--Add Form -->
</template>

<script>
export default {
	data() {
		return {
			addData: {},
			dialogAddVisible: false,
			uid:this.$enum.get("user_info"),
			cTp:this.$enum.get("c_tp"),
			score: [],
			rules: {                    //数据验证规则
				uid: [{ required: true, message: "请输入用户", trigger: "blur" }],
				score: [{ required: true, message: "请输入变动分数", trigger: "blur" }],
			},
		}
	},
	props: {
		refresh: {
			type: Function,
				default: () => {
				},
		}
	},
	created(){
	},
	methods: {
		closed() {
			this.refresh()
		},
		resetForm(formName) {
			this.dialogAddVisible = false;
			this.$refs[formName].resetFields();
		},
		show(){
			this.dialogAddVisible = true;
		},
		handleChooseTool() {
      this.$forceUpdate()
    },
		setScore(pid){
			this.addData.score = ""
			this.score=this.$enum.get("score",pid)
		},
		add(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.$http.post("/score/record", this.addData, {}, true, true)
						.then(res => {
							this.$refs[formName].resetFields()
							this.dialogAddVisible = false
							this.refresh()
						})
						.catch(err => {
							this.$message({
								type: "error",
								message: err.response.data
							});
						})
				} else {
						console.log("error submit!!");
						return false;
				}
			});
		},
	}

}
</script>

<style scoped>
</style>
