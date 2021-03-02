package punishment

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/sql"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/field"
	"github.com/emshop/scoresys/mgrserver/api/modules/db"
)

//PunishmentInfoHandler 惩罚信息处理服务
type PunishmentInfoHandler struct {
}

func NewPunishmentInfoHandler() *PunishmentInfoHandler {
	return &PunishmentInfoHandler{}
}

//PostHandle 添加惩罚信息数据
func (u *PunishmentInfoHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加惩罚信息数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postPunishmentInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	pnID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "惩罚信息"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["pn_id"] = pnID
	count, err := xdb.Execute(sql.InsertPunishmentInfo, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取惩罚信息单条数据
func (u *PunishmentInfoHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取惩罚信息单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getPunishmentInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetPunishmentInfoByPnID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取惩罚信息数据列表
func (u *PunishmentInfoHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取惩罚信息数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryPunishmentInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetPunishmentInfoListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetPunishmentInfoList, m)
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
//PutHandle 更新惩罚信息数据
func (u *PunishmentInfoHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新惩罚信息数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updatePunishmentInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	count,err := hydra.C.DB().GetRegularDB().Execute(sql.UpdatePunishmentInfoByPnID,ctx.Request().GetMap())
	if err != nil||count<1 {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postPunishmentInfoCheckFields = map[string]interface{}{
	field.FieldPnName:"required",
	field.FieldCategory:"required",
	field.FieldScore:"required",
	field.FieldStatus:"required",
	}

var getPunishmentInfoCheckFields = map[string]interface{}{
	field.FieldPnID:"required",
}

var queryPunishmentInfoCheckFields = map[string]interface{}{
	field.FieldPnName:"required",
	field.FieldStatus:"required",
	}

var updatePunishmentInfoCheckFields = map[string]interface{}{
	field.FieldPnName:"required",
	field.FieldCategory:"required",
	field.FieldScore:"required",
	field.FieldStatus:"required",
	}


