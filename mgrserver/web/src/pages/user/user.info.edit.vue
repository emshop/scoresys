<template>
	<el-dialog title="编辑用户信息" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="姓名:" prop="name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.name" placeholder="请输入姓名">
				</el-input>
      </el-form-item>
      
      <el-form-item label="头像:" prop="url">
				<el-input size="medium" maxlength="128"
				clearable v-model="editData.url" placeholder="请输入头像">
				</el-input>
      </el-form-item>
      
      
			<el-form-item prop="birthday" label="出生日期:">
					<el-date-picker size="medium" class="input-cos"  v-model="editData.birthday" type="date" value-format="yyyy-MM-dd"  placeholder="选择日期"></el-date-picker>
			</el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="editData.status" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
    </el-form>
		<div slot="footer" class="dialog-footer">
			<el-button size="medium" @click="dialogFormVisible = false">取 消</el-button>
			<el-button type="success" size="medium" @click="edit">确 定</el-button>
		</div>
	</el-dialog>
</template>

<script>
export default {
	data() {
		return {
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				name: [
					{ required: true, message: "请输入姓名", trigger: "blur" }
				],
				url: [
					{ required: true, message: "请输入头像", trigger: "blur" }
				],
				birthday: [
					{ required: true, message: "请输入出生日期", trigger: "blur" }
				],
				status: [
					{ required: true, message: "请输入状态", trigger: "blur" }
				],
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
		show() {
			this.editData = this.$http.xget("/user/info", { uid: this.editData.uid })
			this.dialogFormVisible = true;
		},
		edit() {
			this.editData.birthday = this.$utility.dateFormat(this.editData.birthday,"yyyy-MM-dd")
			this.$http.put("/user/info", this.editData, {}, true, true)
			.then(res => {			
				this.dialogFormVisible = false;
				this.refresh()
			})
		},
	}
}
</script>

<style scoped>
</style>
