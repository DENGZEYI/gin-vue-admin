import service from '@/utils/request'

// @Tags BusOrder
// @Summary 创建BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusOrder true "创建BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busOrder/createBusOrder [post]
export const createBusOrder = (data) => {
  return service({
    url: '/busOrder/createBusOrder',
    method: 'post',
    data
  })
}

// @Tags BusOrder
// @Summary 删除BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusOrder true "删除BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busOrder/deleteBusOrder [delete]
export const deleteBusOrder = (data) => {
  return service({
    url: '/busOrder/deleteBusOrder',
    method: 'delete',
    data
  })
}

// @Tags BusOrder
// @Summary 删除BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busOrder/deleteBusOrder [delete]
export const deleteBusOrderByIds = (data) => {
  return service({
    url: '/busOrder/deleteBusOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags BusOrder
// @Summary 更新BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusOrder true "更新BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busOrder/updateBusOrder [put]
export const updateBusOrder = (data) => {
  return service({
    url: '/busOrder/updateBusOrder',
    method: 'put',
    data
  })
}

// @Tags BusOrder
// @Summary 用id查询BusOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BusOrder true "用id查询BusOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busOrder/findBusOrder [get]
export const findBusOrder = (params) => {
  return service({
    url: '/busOrder/findBusOrder',
    method: 'get',
    params
  })
}

// @Tags BusOrder
// @Summary 分页获取BusOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BusOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busOrder/getBusOrderList [get]
export const getBusOrderList = (params) => {
  return service({
    url: '/busOrder/getBusOrderList',
    method: 'get',
    params
  })
}
