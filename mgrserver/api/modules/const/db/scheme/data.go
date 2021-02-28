package scheme

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

func init() {
	//注册服务包
	hydra.OnReady(func() error {
		hydra.Installer.DB.AddSQL(
			"INSERT INTO sc_dict_info(id,name,value,type,status,sort_no) VALUES(1,'启用','0','status',0,0),(2,'生活习惯','1','category',0,0);",
		)
		return nil
	})
}
