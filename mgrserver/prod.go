// +build prod

package main

import (
	"embed"
	"fmt"

	_ "github.com/emshop/scoresys/mgrserver/api"
	_ "github.com/emshop/scoresys/mgrserver/api/modules/const/db/scheme"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/processor"
	"github.com/micro-plat/hydra/conf/server/static"
)

//go:embed app/dist
var appweb embed.FS

//go:embed web/dist
var mgrweb embed.FS

func init() {
	//设置配置参数
	hydra.Conf.Web("8090").Header(header.WithCrossDomain()).
		Static(static.WithAutoRewrite(), static.WithEmbed("app/dist", appweb)).
		Processor(processor.WithServicePrefix("api"))

	hydra.Conf.API("8091").Header(header.WithCrossDomain()).
		Static(static.WithAutoRewrite(), static.WithEmbed("web/dist", mgrweb))

	hydra.Conf.Vars().DB().MySQLByConnStr("db", "scoresys:a1qaz2wsxA@tcp(rm-abp1jncd360df9dqvy.mysql.rds.aliyuncs.com:3306)/scoresys?charset=utf8")

	//启动时参数配置检查
	App.OnStarting(func(appConf app.IAPPConf) error {
		if _, err := hydra.C.DB().GetDB(); err != nil {
			return fmt.Errorf("db数据库配置错误,err:%v", err)
		}

		return nil
	})
}
