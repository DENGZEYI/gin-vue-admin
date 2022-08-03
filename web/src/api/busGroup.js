import service from '@/utils/request'

// @Tags BusGroup
// @Summary 创建BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGroup true "创建BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGroup/createBusGroup [post]
export const createBusGroup = (data) => {
  return service({
    url: '/busGroup/createBusGroup',
    method: 'post',
    data
  })
}

// @Tags BusGroup
// @Summary 删除BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGroup true "删除BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGroup/deleteBusGroup [delete]
export const deleteBusGroup = (data) => {
  return service({
    url: '/busGroup/deleteBusGroup',
    method: 'delete',
    data
  })
}

// @Tags BusGroup
// @Summary 删除BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGroup/deleteBusGroup [delete]
export const deleteBusGroupByIds = (data) => {
  return service({
    url: '/busGroup/deleteBusGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags BusGroup
// @Summary 更新BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGroup true "更新BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busGroup/updateBusGroup [put]
export const updateBusGroup = (data) => {
  return service({
    url: '/busGroup/updateBusGroup',
    method: 'put',
    data
  })
}

// @Tags BusGroup
// @Summary 用id查询BusGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BusGroup true "用id查询BusGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busGroup/findBusGroup [get]
export const findBusGroup = (params) => {
  return service({
    url: '/busGroup/findBusGroup',
    method: 'get',
    params
  })
}

// @Tags BusGroup
// @Summary 分页获取BusGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BusGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGroup/getBusGroupList [get]
export const getBusGroupList = (params) => {
  return service({
    url: '/busGroup/getBusGroupList',
    method: 'get',
    params
  })
}
