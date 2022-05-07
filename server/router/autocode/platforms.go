package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlatformsRouter struct {
}

// InitPlatformsRouter 初始化 Platforms 路由信息
func (s *PlatformsRouter) InitPlatformsRouter(Router *gin.RouterGroup) {
	platformsRouter := Router.Group("platforms").Use(middleware.OperationRecord())
	platformsRouterWithoutRecord := Router.Group("platforms")
	var platformsApi = v1.ApiGroupApp.AutoCodeApiGroup.PlatformsApi
	{
		platformsRouter.POST("createPlatforms", platformsApi.CreatePlatforms)   // 新建Platforms
		platformsRouter.DELETE("deletePlatforms", platformsApi.DeletePlatforms) // 删除Platforms
		platformsRouter.DELETE("deletePlatformsByIds", platformsApi.DeletePlatformsByIds) // 批量删除Platforms
		platformsRouter.PUT("updatePlatforms", platformsApi.UpdatePlatforms)    // 更新Platforms
	}
	{
		platformsRouterWithoutRecord.GET("findPlatforms", platformsApi.FindPlatforms)        // 根据ID获取Platforms
		platformsRouterWithoutRecord.GET("getPlatformsList", platformsApi.GetPlatformsList)  // 获取Platforms列表
	}
}
