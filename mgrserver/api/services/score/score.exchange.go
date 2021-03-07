package score

import (
	"net/http"

	"github.com/emshop/scoresys/mgrserver/api/modules/const/field"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/sql"
	"github.com/emshop/scoresys/mgrserver/api/modules/db"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//ExchangeHandle 添加分数
func (u *ScoreRecordHandler) ExchangeHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加分数--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(exchangeScoreRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	ctx.Log().Info("3.获取批次编号")
	batchID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "批次编号"})
	if err != nil {
		xdb.Rollback()
		return err
	}

	ctx.Log().Info("4.保存数据")
	//获取添加记录
	rwID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "兑换记录"})
	if err != nil {
		xdb.Rollback()
		return err
	}

	score := ctx.Request().GetInt(field.FieldScore)
	if score > 0 {
		score = -1 * score
	}

	input := ctx.Request().GetMap()
	input[field.FieldRwpuID] = 0
	input[field.FieldCTp] = 3
	input[field.FieldRcID] = rwID
	input[field.FieldBatchID] = batchID
	input[field.FieldScore] = score

	input[field.FieldContent] = ctx.Request().GetString(field.FieldContent)
	count, err := xdb.Execute(sql.InsertScoreRecord, input)
	if err != nil || count < 1 {
		xdb.Rollback()
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("5.更新用户分数")
	count, err = xdb.Execute(sql.UpdateUserScoreByUid, input)
	if err != nil || count < 1 {
		xdb.Rollback()
		return errs.NewErrorf(http.StatusNotExtended, "更新用户分数出错:%+v", err)
	}
	xdb.Commit()
	ctx.Log().Info("3.返回结果")
	return "success"
}

var exchangeScoreRecordCheckFields = map[string]interface{}{
	field.FieldScore:   "required",
	field.FieldContent: "required",
	field.FieldUid:     "required",
}
