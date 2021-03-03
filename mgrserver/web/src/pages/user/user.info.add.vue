<template>
  <!-- Add Form -->
  <el-dialog title="添加用户信息" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="姓名:" prop="name">
				<el-input size="medium" maxlength="64"
				 clearable v-model="addData.name" placeholder="请输入姓名">
				</el-input>
      </el-form-item>
      
      <el-form-item label="头像:" prop="url">
				<el-input size="medium" maxlength="128"
				 clearable v-model="addData.url" placeholder="请输入头像">
				</el-input>
      </el-form-item>
      
      
			<el-form-item prop="birthday" label="出生日期:">
					<el-date-picker size="medium" class="input-cos"  v-model="addData.birthday" type="date" value-format="yyyy-MM-dd"  placeholder="选择日期"></el-date-picker>
			</el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="addData.status"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
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
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
				url: [{ required: true, message: "请输入头像", trigger: "blur" }],
				birthday: [{ required: true, message: "请输入出生日期", trigger: "blur" }],
				status: [{ required: true, message: "请输入状态", trigger: "blur" }],
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
		add(formName) {
			this.addData.birthday = this.$utility.dateFormat(this.addData.birthday,"yyyy-MM-dd")
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.$http.post("/user/info", this.addData, {}, true, true)
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
