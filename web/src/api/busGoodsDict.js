import service from '@/utils/request'

// @Tags busGoodsDict
// @Summary 创建busGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.busGoodsDict true "创建busGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGoodsDict/createbusGoodsDict [post]
export const createBusGoodsDict = (data) => {
  return service({
    url: '/busGoodsDict/createBusGoodsDict',
    method: 'post',
    data
  })
}

// @Tags busGoodsDict
// @Summary 删除busGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.busGoodsDict true "删除busGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGoodsDict/deletebusGoodsDict [delete]
export const deleteBusGoodsDict = (data) => {
  return service({
    url: '/busGoodsDict/deleteBusGoodsDict',
    method: 'delete',
    data
  })
}

// @Tags busGoodsDict
// @Summary 删除busGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除busGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGoodsDict/deletebusGoodsDict [delete]
export const deleteBusGoodsDictByIds = (data) => {
  return service({
    url: '/busGoodsDict/deleteBusGoodsDictByIds',
    method: 'delete',
    data
  })
}

export const applyBusGoods = (data) => {
  return service({
    url: '/busGoodsDict/applyBusGoods',
    method: 'post',
    data
  })
}

// @Tags busGoodsDict
// @Summary 更新busGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.busGoodsDict true "更新busGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busGoodsDict/updatebusGoodsDict [put]
export const updateBusGoodsDict = (data) => {
  return service({
    url: '/busGoodsDict/updateBusGoodsDict',
    method: 'put',
    data
  })
}

// @Tags busGoodsDict
// @Summary 用id查询busGoodsDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.busGoodsDict true "用id查询busGoodsDict"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busGoodsDict/findbusGoodsDict [get]
export const findBusGoodsDict = (params) => {
  return service({
    url: '/busGoodsDict/findBusGoodsDict',
    method: 'get',
    params
  })
}

// @Tags busGoodsDict
// @Summary 分页获取busGoodsDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取busGoodsDict列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGoodsDict/getbusGoodsDictList [get]
export const getBusGoodsDictList = (params) => {
  return service({
    url: '/busGoodsDict/getBusGoodsDictList',
    method: 'get',
    params
  })
}
