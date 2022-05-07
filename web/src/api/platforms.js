import service from '@/utils/request'

// @Tags Platforms
// @Summary 创建Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platforms true "创建Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platforms/createPlatforms [post]
export const createPlatforms = (data) => {
  return service({
    url: '/platforms/createPlatforms',
    method: 'post',
    data
  })
}

// @Tags Platforms
// @Summary 删除Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platforms true "删除Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platforms/deletePlatforms [delete]
export const deletePlatforms = (data) => {
  return service({
    url: '/platforms/deletePlatforms',
    method: 'delete',
    data
  })
}

// @Tags Platforms
// @Summary 删除Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platforms/deletePlatforms [delete]
export const deletePlatformsByIds = (data) => {
  return service({
    url: '/platforms/deletePlatformsByIds',
    method: 'delete',
    data
  })
}

// @Tags Platforms
// @Summary 更新Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Platforms true "更新Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platforms/updatePlatforms [put]
export const updatePlatforms = (data) => {
  return service({
    url: '/platforms/updatePlatforms',
    method: 'put',
    data
  })
}

// @Tags Platforms
// @Summary 用id查询Platforms
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Platforms true "用id查询Platforms"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platforms/findPlatforms [get]
export const findPlatforms = (params) => {
  return service({
    url: '/platforms/findPlatforms',
    method: 'get',
    params
  })
}

// @Tags Platforms
// @Summary 分页获取Platforms列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Platforms列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platforms/getPlatformsList [get]
export const getPlatformsList = (params) => {
  return service({
    url: '/platforms/getPlatformsList',
    method: 'get',
    params
  })
}
