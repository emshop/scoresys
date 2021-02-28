
package scheme

import (
	"github.com/micro-plat/hydra"
	_ "github.com/go-sql-driver/mysql"
)
		
func init() {
	//注册服务包
	hydra.OnReadying(func() error {
		hydra.Installer.DB.AddSQL(
		sc_reward_info,
		sc_punishment_info,
		sc_gif_info,
		sc_user_info,
		sc_score_record,
		sc_dict_info,
		SEQ_IDS,
		)
		return nil
	}) 
}
