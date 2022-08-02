import service from '@/utils/request'

// @Tags TestDic
// @Summary 创建TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestDic true "创建TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDic/createTestDic [post]
export const createTestDic = (data) => {
  return service({
    url: '/testDic/createTestDic',
    method: 'post',
    data
  })
}

// @Tags TestDic
// @Summary 删除TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestDic true "删除TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testDic/deleteTestDic [delete]
export const deleteTestDic = (data) => {
  return service({
    url: '/testDic/deleteTestDic',
    method: 'delete',
    data
  })
}

// @Tags TestDic
// @Summary 删除TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testDic/deleteTestDic [delete]
export const deleteTestDicByIds = (data) => {
  return service({
    url: '/testDic/deleteTestDicByIds',
    method: 'delete',
    data
  })
}

// @Tags TestDic
// @Summary 更新TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestDic true "更新TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testDic/updateTestDic [put]
export const updateTestDic = (data) => {
  return service({
    url: '/testDic/updateTestDic',
    method: 'put',
    data
  })
}

// @Tags TestDic
// @Summary 用id查询TestDic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestDic true "用id查询TestDic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testDic/findTestDic [get]
export const findTestDic = (params) => {
  return service({
    url: '/testDic/findTestDic',
    method: 'get',
    params
  })
}

// @Tags TestDic
// @Summary 分页获取TestDic列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestDic列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testDic/getTestDicList [get]
export const getTestDicList = (params) => {
  return service({
    url: '/testDic/getTestDicList',
    method: 'get',
    params
  })
}
