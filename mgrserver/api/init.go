package api

import (
	"github.com/micro-plat/hydra"
	"github.com/emshop/scoresys/mgrserver/api/services/dict"
	"github.com/emshop/scoresys/mgrserver/api/services/gif"
	"github.com/emshop/scoresys/mgrserver/api/services/punishment"
	"github.com/emshop/scoresys/mgrserver/api/services/reward"
	"github.com/emshop/scoresys/mgrserver/api/services/score"
	"github.com/emshop/scoresys/mgrserver/api/services/system"
	"github.com/emshop/scoresys/mgrserver/api/services/user"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	hydra.OnReady(func() {
		hydra.S.Web("/system/enums", system.NewSystemEnumsHandler())
		hydra.S.Web("/dict/info", dict.NewDictInfoHandler())
		hydra.S.Web("/gif/info", gif.NewGifInfoHandler())
		hydra.S.Web("/punishment/info", punishment.NewPunishmentInfoHandler())
		hydra.S.Web("/reward/info", reward.NewRewardInfoHandler())
		hydra.S.Web("/score/record", score.NewScoreRecordHandler())
		hydra.S.Web("/user/info", user.NewUserInfoHandler())
	})
}

