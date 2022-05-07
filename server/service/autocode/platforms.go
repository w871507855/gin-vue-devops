package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type PlatformsService struct {
}

// CreatePlatforms 创建Platforms记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformsService *PlatformsService) CreatePlatforms(platforms autocode.Platforms) (err error) {
	err = global.GVA_DB.Create(&platforms).Error
	return err
}

// DeletePlatforms 删除Platforms记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformsService *PlatformsService)DeletePlatforms(platforms autocode.Platforms) (err error) {
	err = global.GVA_DB.Delete(&platforms).Error
	return err
}

// DeletePlatformsByIds 批量删除Platforms记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformsService *PlatformsService)DeletePlatformsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Platforms{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePlatforms 更新Platforms记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformsService *PlatformsService)UpdatePlatforms(platforms autocode.Platforms) (err error) {
	err = global.GVA_DB.Save(&platforms).Error
	return err
}

// GetPlatforms 根据id获取Platforms记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformsService *PlatformsService)GetPlatforms(id uint) (err error, platforms autocode.Platforms) {
	err = global.GVA_DB.Where("id = ?", id).First(&platforms).Error
	return
}

// GetPlatformsInfoList 分页获取Platforms记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformsService *PlatformsService)GetPlatformsInfoList(info autoCodeReq.PlatformsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Platforms{})
    var platformss []autocode.Platforms
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&platformss).Error
	return err, platformss, total
}
