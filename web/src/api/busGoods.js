import service from '@/utils/request'

// @Tags BusGoods
// @Summary 创建BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGoods true "创建BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGoods/createBusGoods [post]
export const createBusGoods = (data) => {
  return service({
    url: '/busGoods/createBusGoods',
    method: 'post',
    data
  })
}

// @Tags BusGoods
// @Summary 删除BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGoods true "删除BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGoods/deleteBusGoods [delete]
export const deleteBusGoods = (data) => {
  return service({
    url: '/busGoods/deleteBusGoods',
    method: 'delete',
    data
  })
}

// @Tags BusGoods
// @Summary 删除BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGoods/deleteBusGoods [delete]
export const deleteBusGoodsByIds = (data) => {
  return service({
    url: '/busGoods/deleteBusGoodsByIds',
    method: 'delete',
    data
  })
}

// @Tags BusGoods
// @Summary 更新BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGoods true "更新BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busGoods/updateBusGoods [put]
export const updateBusGoods = (data) => {
  return service({
    url: '/busGoods/updateBusGoods',
    method: 'put',
    data
  })
}

// @Tags BusGoods
// @Summary 用id查询BusGoods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BusGoods true "用id查询BusGoods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busGoods/findBusGoods [get]
export const findBusGoods = (params) => {
  return service({
    url: '/busGoods/findBusGoods',
    method: 'get',
    params
  })
}

// @Tags BusGoods
// @Summary 分页获取BusGoods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BusGoods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGoods/getBusGoodsList [get]
export const getBusGoodsList = (params) => {
  return service({
    url: '/busGoods/getBusGoodsList',
    method: 'get',
    params
  })
}
