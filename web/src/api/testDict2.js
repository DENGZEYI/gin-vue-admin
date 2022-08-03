import service from '@/utils/request'

// @Tags TestDict2
// @Summary 创建TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestDict2 true "创建TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDict2/createTestDict2 [post]
export const createTestDict2 = (data) => {
  return service({
    url: '/testDict2/createTestDict2',
    method: 'post',
    data
  })
}

// @Tags TestDict2
// @Summary 删除TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestDict2 true "删除TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testDict2/deleteTestDict2 [delete]
export const deleteTestDict2 = (data) => {
  return service({
    url: '/testDict2/deleteTestDict2',
    method: 'delete',
    data
  })
}

// @Tags TestDict2
// @Summary 删除TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testDict2/deleteTestDict2 [delete]
export const deleteTestDict2ByIds = (data) => {
  return service({
    url: '/testDict2/deleteTestDict2ByIds',
    method: 'delete',
    data
  })
}

// @Tags TestDict2
// @Summary 更新TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestDict2 true "更新TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testDict2/updateTestDict2 [put]
export const updateTestDict2 = (data) => {
  return service({
    url: '/testDict2/updateTestDict2',
    method: 'put',
    data
  })
}

// @Tags TestDict2
// @Summary 用id查询TestDict2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestDict2 true "用id查询TestDict2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testDict2/findTestDict2 [get]
export const findTestDict2 = (params) => {
  return service({
    url: '/testDict2/findTestDict2',
    method: 'get',
    params
  })
}

// @Tags TestDict2
// @Summary 分页获取TestDict2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestDict2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDict2/getTestDict2List [get]
export const getTestDict2List = (params) => {
  return service({
    url: '/testDict2/getTestDict2List',
    method: 'get',
    params
  })
}
