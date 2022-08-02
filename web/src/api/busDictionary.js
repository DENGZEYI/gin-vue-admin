import service from '@/utils/request'

// @Tags SysDictionary
// @Summary 查询应用层字典(组、商品、供应商)
// @Security ApiKeyAuth
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busDictionary/findBusDictionary [get]
export const findBusDictionary = (params) => {
    return service({
      url: '/busDictionary/findBusDictionary',
      method: 'get',
      params
    })
}