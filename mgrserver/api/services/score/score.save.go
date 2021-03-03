package score

import (
	"fmt"
	"net/http"

	"github.com/emshop/scoresys/mgrserver/api/modules/const/field"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/sql"
	"github.com/emshop/scoresys/mgrserver/api/modules/db"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//AddHandle 添加分数
func (u *ScoreRecordHandler) AddHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加分数--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(addScoreRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//更新用户分数
	count, err := xdb.Execute(sql.UpdateUserScoreByUid, ctx.Request().GetMap())
	if err != nil || count < 1 {
		xdb.Rollback()
		return errs.NewErrorf(http.StatusNotExtended, "更新用户分数出错:%+v", err)
	}

	//添加记录
	rwID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "分数记录"})
	if err != nil {
		xdb.Rollback()
		return err
	}

	xmp, err := xdb.Query(sql.GetUserInfoByUid, ctx.Request().GetMap())
	if err != nil {
		xdb.Rollback()
		return err
	}
	if xmp.Len() == 0 {
		xdb.Rollback()
		return fmt.Errorf("未找到用户%s", ctx.Request().GetString(field.FieldUid))
	}

	//保存分数信息
	input := ctx.Request().GetMap()
	input[field.FieldRcID] = rwID
	input[field.FieldRemain] = xmp.Get(0).GetString(field.FieldScore)
	count, err = xdb.Execute(sql.InsertScoreRecord, input)
	if err != nil || count < 1 {
		xdb.Rollback()
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}
	xdb.Commit()
	ctx.Log().Info("3.返回结果")
	return "success"
}

var addScoreRecordCheckFields = map[string]interface{}{
	field.FieldScore: "required",
	field.FieldUid:   "required",
}
