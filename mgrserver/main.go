package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("emshop", "平台中文名"),
	hydra.WithSystemName("scoresys", "积分管理"),
	hydra.WithServerTypes(http.Web),
)

func main() {
	App.Start()
}
