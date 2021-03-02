package user

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/sql"
	"github.com/emshop/scoresys/mgrserver/api/modules/const/field"
	"github.com/emshop/scoresys/mgrserver/api/modules/db"
)

//UserInfoHandler 用户信息处理服务
type UserInfoHandler struct {
}

func NewUserInfoHandler() *UserInfoHandler {
	return &UserInfoHandler{}
}

//PostHandle 添加用户信息数据
func (u *UserInfoHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加用户信息数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postUserInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	uid, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "用户信息"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["uid"] = uid
	count, err := xdb.Execute(sql.InsertUserInfo, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取用户信息单条数据
func (u *UserInfoHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取用户信息单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getUserInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetUserInfoByUid,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取用户信息数据列表
func (u *UserInfoHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取用户信息数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryUserInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetUserInfoListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetUserInfoList, m)
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
//PutHandle 更新用户信息数据
func (u *UserInfoHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新用户信息数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateUserInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	count,err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateUserInfoByUid,ctx.Request().GetMap())
	if err != nil||count<1 {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postUserInfoCheckFields = map[string]interface{}{
	field.FieldName:"required",
	field.FieldURL:"required",
	field.FieldStatus:"required",
	}

var getUserInfoCheckFields = map[string]interface{}{
	field.FieldUid:"required",
}

var queryUserInfoCheckFields = map[string]interface{}{
	field.FieldName:"required",
	field.FieldStatus:"required",
	}

var updateUserInfoCheckFields = map[string]interface{}{
	field.FieldName:"required",
	field.FieldURL:"required",
	field.FieldStatus:"required",
	}


