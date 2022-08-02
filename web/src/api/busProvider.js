import service from '@/utils/request'

// @Tags BusProvider
// @Summary 创建BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusProvider true "创建BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busProvider/createBusProvider [post]
export const createBusProvider = (data) => {
  return service({
    url: '/busProvider/createBusProvider',
    method: 'post',
    data
  })
}

// @Tags BusProvider
// @Summary 删除BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusProvider true "删除BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busProvider/deleteBusProvider [delete]
export const deleteBusProvider = (data) => {
  return service({
    url: '/busProvider/deleteBusProvider',
    method: 'delete',
    data
  })
}

// @Tags BusProvider
// @Summary 删除BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busProvider/deleteBusProvider [delete]
export const deleteBusProviderByIds = (data) => {
  return service({
    url: '/busProvider/deleteBusProviderByIds',
    method: 'delete',
    data
  })
}

// @Tags BusProvider
// @Summary 更新BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusProvider true "更新BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busProvider/updateBusProvider [put]
export const updateBusProvider = (data) => {
  return service({
    url: '/busProvider/updateBusProvider',
    method: 'put',
    data
  })
}

// @Tags BusProvider
// @Summary 用id查询BusProvider
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BusProvider true "用id查询BusProvider"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busProvider/findBusProvider [get]
export const findBusProvider = (params) => {
  return service({
    url: '/busProvider/findBusProvider',
    method: 'get',
    params
  })
}

// @Tags BusProvider
// @Summary 分页获取BusProvider列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BusProvider列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busProvider/getBusProviderList [get]
export const getBusProviderList = (params) => {
  return service({
    url: '/busProvider/getBusProviderList',
    method: 'get',
    params
  })
}
