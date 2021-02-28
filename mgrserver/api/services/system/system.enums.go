package system

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//SystemEnumsHandler 枚举数据查询服务
type SystemEnumsHandler struct {
}

//NewSystemEnumsHandler 枚举数据查询服务
func NewSystemEnumsHandler() *SystemEnumsHandler {
	return &SystemEnumsHandler{}
}

//QueryHandle 枚举数据查询服务
func (o *SystemEnumsHandler) QueryHandle(ctx hydra.IContext) interface{} {

	//根据传入的枚举类型获取数据
	tp := ctx.Request().GetString("dic_type")
	if tp != "" && enumsMap[tp] != "" {

		items, err := hydra.C.DB().GetRegularDB().Query(enumsMap[tp], ctx.Request().GetMap())
		if err != nil {
			return err
		}
		return items
	}

	//查询所有枚举数据
	list := types.XMaps{}
	for _, sql := range enumsMap {
		items, err := hydra.C.DB().GetRegularDB().Query(sql, ctx.Request().GetMap())
		if err != nil {
			return err
		}
		list = append(list, items...)
	}
	return list
}

var enumsMap = map[string]string{
	"reward_info":     `select 'reward_info' type , t.rw_id value , t.rw_name name  from sc_reward_info t `,
	"punishment_info": `select 'punishment_info' type , t.pn_name name , t.score value  from sc_punishment_info t `,
	"gif_info":        `select 'gif_info' type , t.gif_name name , t.score value  from sc_gif_info t `,
	"user_info":       `select 'user_info' type , t.uid value , t.name name   from sc_user_info t `,
	"score_record":    `select 'score_record' type , t.rc_id value , t.uid name   from sc_score_record t `,
	"dict_info":       `select t.type type , t.value value , t.name name  from sc_dict_info t `,
}
