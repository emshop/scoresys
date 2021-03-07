package score

import (
	"net/http"

	"github.com/emshop/scoresys/mgrserver/api/modules/const/field"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/sql"
	"github.com/emshop/scoresys/mgrserver/api/modules/db"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
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
	ctx.Log().Info("3.获取批次编号")
	batchID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "批次编号"})
	if err != nil {
		xdb.Rollback()
		return err
	}

	ctx.Log().Info("4.循环保存数据")
	ids := ctx.Request().GetStrings(field.FieldIDs)
	for _, vid := range ids {

		score := 0
		content := ""
		ctp := 0
		data, err := xdb.Query(sql.GetRewardInfoByRwID, map[string]interface{}{field.FieldRwID: vid})
		if err != nil {
			xdb.Rollback()
			return err
		}
		if data.Len() > 0 {
			score = data.Get(0).GetInt(field.FieldScore)
			content = data.Get(0).GetString(field.FieldRwName)
		} else {
			ctp = 1
			data, err := xdb.Query(sql.GetPunishmentInfoByPnID, map[string]interface{}{field.FieldPnID: vid})
			if err != nil {
				xdb.Rollback()
				return err
			}
			score = data.Get(0).GetInt(field.FieldScore)
			content = data.Get(0).GetString(field.FieldPnName)
			if score > 0 {
				score *= -1
			}
		}

		//获取添加记录
		rwID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "分数记录"})
		if err != nil {
			xdb.Rollback()
			return err
		}

		input := ctx.Request().GetMap()
		input[field.FieldRwpuID] = vid
		input[field.FieldCTp] = ctp
		input[field.FieldRcID] = rwID
		input[field.FieldScore] = score
		input[field.FieldBatchID] = batchID
		input[field.FieldContent] = content
		count, err := xdb.Execute(sql.InsertScoreRecord, input)
		if err != nil || count < 1 {
			xdb.Rollback()
			return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
		}
	}
	ctx.Log().Info("5.更新用户分数")
	input := ctx.Request().GetMap()
	input[field.FieldBatchID] = batchID
	count, err := xdb.Execute(sql.UpdateUserScoreByUid, input)
	if err != nil || count < 1 {
		xdb.Rollback()
		return errs.NewErrorf(http.StatusNotExtended, "更新用户分数出错:%+v", err)
	}
	xdb.Commit()
	ctx.Log().Info("3.返回结果")
	return "success"
}

var addScoreRecordCheckFields = map[string]interface{}{
	field.FieldIDs: "required",
	field.FieldUid: "required",
}

//QueryHandle  获取分数记录数据列表
func (u *ScoreRecordHandler) QHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取分数记录数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryScoreRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(GetScoreRecordListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}

	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(GetScoreRecordList2, m)
		if err != nil {
			return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
		}
	}

	ctx.Log().Info("3.返回结果")
	return map[string]interface{}{
		"items": items,
		"count": types.GetInt(count),
	}
}

//GetScoreRecordList 查询分数记录列表数据
const GetScoreRecordList2 = `
select
	t.c_tp,
	t.uid,
	t.rwpu_id,
	t.score,
	t.content,
	t.create_time 
from sc_score_record t
where
	&t.uid
	and t.create_time >= @create_time 
order by t.create_time desc
limit @ps offset @offset
`

//GetScoreRecordListCount 获取分数记录列表条数
const GetScoreRecordListCount = `
select count(1)
from sc_score_record t
where
	&t.uid
	and t.create_time >= @create_time`
