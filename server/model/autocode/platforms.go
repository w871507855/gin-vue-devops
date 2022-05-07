// 自动生成模板Platforms
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Platforms 结构体
// 如果含有time.Time 请自行import time包
type Platforms struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:;"`
      Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:;"`
}


