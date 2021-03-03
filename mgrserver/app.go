// +build app

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
var fs embed.FS

func init() {
	//设置配置参数
	hydra.Conf.Web("8080").Header(header.WithCrossDomain()).
		Static(static.WithAutoRewrite(), static.WithEmbed("app/dist", fs)).
		Processor(processor.WithServicePrefix("api"))
	// hydra.Conf.Vars().DB().MySQL("db", "hydra", "123456", "192.168.0.36:3306", "hydra", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(222.209.84.37:10036)/hydra?charset=utf8")

	//启动时参数配置检查
	App.OnStarting(func(appConf app.IAPPConf) error {
		if _, err := hydra.C.DB().GetDB(); err != nil {
			return fmt.Errorf("db数据库配置错误,err:%v", err)
		}

		return nil
	})
}
