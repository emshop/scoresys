package reward

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/sql"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	"github.com/emshop/scoresys/mgrserver/api/modules/db"
)

//RewardInfoHandler 奖励信息处理服务
type RewardInfoHandler struct {
}

func NewRewardInfoHandler() *RewardInfoHandler {
	return &RewardInfoHandler{}
}

//PostHandle 添加奖励信息数据
func (u *RewardInfoHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加奖励信息数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postRewardInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	rwID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "奖励信息"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["rw_id"] = rwID
	count, err := xdb.Execute(sql.InsertRewardInfo, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取奖励信息单条数据
func (u *RewardInfoHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取奖励信息单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getRewardInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetRewardInfoByRwID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取奖励信息数据列表
func (u *RewardInfoHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取奖励信息数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryRewardInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetRewardInfoListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetRewardInfoList, m)
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
//PutHandle 更新奖励信息数据
func (u *RewardInfoHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新奖励信息数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateRewardInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	count,err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateRewardInfoByRwID,ctx.Request().GetMap())
	if err != nil||count<1 {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postRewardInfoCheckFields = map[string]interface{}{
	field.FieldRwName:"required",
	field.FieldCategory:"required",
	field.FieldScore:"required",
	field.FieldStatus:"required",
	}

var getRewardInfoCheckFields = map[string]interface{}{
	field.FieldRwID:"required",
}

var queryRewardInfoCheckFields = map[string]interface{}{
	field.FieldRwName:"required",
	field.FieldCategory:"required",
	field.FieldStatus:"required",
	}

var updateRewardInfoCheckFields = map[string]interface{}{
	field.FieldRwName:"required",
	field.FieldCategory:"required",
	field.FieldScore:"required",
	field.FieldStatus:"required",
	}


