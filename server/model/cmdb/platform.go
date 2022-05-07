package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Platform struct {
	global.GVA_MODEL
	// 名称
	Name string `json:"name" gorm:"comment:名称"`
	// 系统类型
	Type string `json:"type" gorm:"comment:系统类型"`
	// 备注
	Comment string `json:"comment" gorm:"comment:备注"`
}
