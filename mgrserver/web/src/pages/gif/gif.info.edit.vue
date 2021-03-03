<template>
	<el-dialog title="编辑礼品信息" width="25%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="礼品名称:" prop="gif_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.gif_name" placeholder="请输入礼品名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="分数:" prop="score">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.score" placeholder="请输入分数">
				</el-input>
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
				gif_name: [
					{ required: true, message: "请输入礼品名称", trigger: "blur" }
				],
				score: [
					{ required: true, message: "请输入分数", trigger: "blur" }
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
			this.editData = this.$http.xget("/gif/info", { gif_id: this.editData.gif_id })
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/gif/info", this.editData, {}, true, true)
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
