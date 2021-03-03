<template>
	<el-dialog title="编辑分数记录" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      
			<el-form-item label="类型:" prop="c_tp">
				<el-select size="medium" style="width: 100%;"	v-model="editData.c_tp" clearable filterable class="input-cos" placeholder="---请选择---" @change="setScore(editData.c_tp)"	>
					<el-option v-for="(item, index) in cTp" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="变动分数:" prop="score">
				<el-select size="medium" style="width: 100%;"	v-model="editData.score" clearable filterable class="input-cos" placeholder="---请选择---" @change="handleChooseTool()"	>
					<el-option v-for="(item, index) in score" :key="index" :value="item.value" :label="item.name"></el-option>
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
      cTp: this.$enum.get("c_tp"),
      score: [],
			rules: {                    //数据验证规则
				score: [
					{ required: true, message: "请输入变动分数", trigger: "blur" }
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
		handleChooseTool() {
      this.$forceUpdate()
    },
		show() {
			this.editData = this.$http.xget("/score/record", { rc_id: this.editData.rc_id })
			this.dialogFormVisible = true;
		},
		setScore(pid){
			this.editData.score = ""
			this.score=this.$enum.get("score",pid)
		},
		edit() {
			this.$http.put("/score/record", this.editData, {}, true, true)
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
