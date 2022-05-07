package cmdb

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Asset struct {
	global.GVA_MODEL
	// 主机名
	Hostname string `json:"hostname" gorm:"comment:主机名"`
	// 服务器ip地址
	IP string `json:"ip" gorm:"comment:ip地址"`
	// 协议组
	Protocol  string `json:"protocol" gorm:"comment:协议组"`
	// 系统平台

	// 公网ip

	// 网域

	// 特权用户

	// 节点

	// 标签

	// 备注
}