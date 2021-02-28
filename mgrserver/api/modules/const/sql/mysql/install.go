package mysql

import "github.com/micro-plat/hydra"

func init() {
	hydra.Installer.DB.AddSQL(sys_sequence_info)
}

