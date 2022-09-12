import service from '@/utils/request'

// @Tags BusGroup
// @Summary 分页获取BusIngress列表
export const getBusIngressList = (params) => {
    return service({
        url: '/busIngress/getBusIngressList',
        method: 'get',
        params
    })
}