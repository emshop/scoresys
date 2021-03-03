package score

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/sql"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	
)

//ScoreRecordHandler 分数记录处理服务
type ScoreRecordHandler struct {
}

func NewScoreRecordHandler() *ScoreRecordHandler {
	return &ScoreRecordHandler{}
}

//PostHandle 添加分数记录数据
func (u *ScoreRecordHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加分数记录数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postScoreRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	count, err := hydra.C.DB().GetRegularDB().Execute(sql.InsertScoreRecord,ctx.Request().GetMap())
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取分数记录单条数据
func (u *ScoreRecordHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取分数记录单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getScoreRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetScoreRecordByRcID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取分数记录数据列表
func (u *ScoreRecordHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取分数记录数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryScoreRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetScoreRecordListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetScoreRecordList, m)
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
//PutHandle 更新分数记录数据
func (u *ScoreRecordHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新分数记录数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateScoreRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	count,err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateScoreRecordByRcID,ctx.Request().GetMap())
	if err != nil||count<1 {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postScoreRecordCheckFields = map[string]interface{}{
	field.FieldUid:"required",
	field.FieldCTp:"required",
	field.FieldScore:"required",
	}

var getScoreRecordCheckFields = map[string]interface{}{
	field.FieldRcID:"required",
}

var queryScoreRecordCheckFields = map[string]interface{}{
	field.FieldUid:"required",
	field.FieldCTp:"required",
	}

var updateScoreRecordCheckFields = map[string]interface{}{
	field.FieldCTp:"required",
	field.FieldScore:"required",
	}


