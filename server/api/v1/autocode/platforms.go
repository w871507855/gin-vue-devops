package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PlatformsApi struct {
}

var platformsService = service.ServiceGroupApp.AutoCodeServiceGroup.PlatformsService


// CreatePlatforms 创建Platforms
// @Tags Platforms
// @Summary 创建Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Platforms true "创建Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platforms/createPlatforms [post]
func (platformsApi *PlatformsApi) CreatePlatforms(c *gin.Context) {
	var platforms autocode.Platforms
	_ = c.ShouldBindJSON(&platforms)
	if err := platformsService.CreatePlatforms(platforms); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlatforms 删除Platforms
// @Tags Platforms
// @Summary 删除Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Platforms true "删除Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platforms/deletePlatforms [delete]
func (platformsApi *PlatformsApi) DeletePlatforms(c *gin.Context) {
	var platforms autocode.Platforms
	_ = c.ShouldBindJSON(&platforms)
	if err := platformsService.DeletePlatforms(platforms); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlatformsByIds 批量删除Platforms
// @Tags Platforms
// @Summary 批量删除Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /platforms/deletePlatformsByIds [delete]
func (platformsApi *PlatformsApi) DeletePlatformsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := platformsService.DeletePlatformsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlatforms 更新Platforms
// @Tags Platforms
// @Summary 更新Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Platforms true "更新Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platforms/updatePlatforms [put]
func (platformsApi *PlatformsApi) UpdatePlatforms(c *gin.Context) {
	var platforms autocode.Platforms
	_ = c.ShouldBindJSON(&platforms)
	if err := platformsService.UpdatePlatforms(platforms); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlatforms 用id查询Platforms
// @Tags Platforms
// @Summary 用id查询Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Platforms true "用id查询Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platforms/findPlatforms [get]
func (platformsApi *PlatformsApi) FindPlatforms(c *gin.Context) {
	var platforms autocode.Platforms
	_ = c.ShouldBindQuery(&platforms)
	if err, replatforms := platformsService.GetPlatforms(platforms.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replatforms": replatforms}, c)
	}
}

// GetPlatformsList 分页获取Platforms列表
// @Tags Platforms
// @Summary 分页获取Platforms列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.PlatformsSearch true "分页获取Platforms列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platforms/getPlatformsList [get]
func (platformsApi *PlatformsApi) GetPlatformsList(c *gin.Context) {
	var pageInfo autocodeReq.PlatformsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := platformsService.GetPlatformsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
