import service from '@/utils/request'

// @Tags BusFactory
// @Summary 创建BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusFactory true "创建BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busFactory/createBusFactory [post]
export const createBusFactory = (data) => {
  return service({
    url: '/busFactory/createBusFactory',
    method: 'post',
    data
  })
}

// @Tags BusFactory
// @Summary 删除BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusFactory true "删除BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busFactory/deleteBusFactory [delete]
export const deleteBusFactory = (data) => {
  return service({
    url: '/busFactory/deleteBusFactory',
    method: 'delete',
    data
  })
}

// @Tags BusFactory
// @Summary 删除BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busFactory/deleteBusFactory [delete]
export const deleteBusFactoryByIds = (data) => {
  return service({
    url: '/busFactory/deleteBusFactoryByIds',
    method: 'delete',
    data
  })
}

// @Tags BusFactory
// @Summary 更新BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusFactory true "更新BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busFactory/updateBusFactory [put]
export const updateBusFactory = (data) => {
  return service({
    url: '/busFactory/updateBusFactory',
    method: 'put',
    data
  })
}

// @Tags BusFactory
// @Summary 用id查询BusFactory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BusFactory true "用id查询BusFactory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busFactory/findBusFactory [get]
export const findBusFactory = (params) => {
  return service({
    url: '/busFactory/findBusFactory',
    method: 'get',
    params
  })
}

// @Tags BusFactory
// @Summary 分页获取BusFactory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BusFactory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busFactory/getBusFactoryList [get]
export const getBusFactoryList = (params) => {
  return service({
    url: '/busFactory/getBusFactoryList',
    method: 'get',
    params
  })
}
