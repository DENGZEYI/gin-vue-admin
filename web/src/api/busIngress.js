import service from '@/utils/request'


// @Summary 分页获取BusIngress列表
export const getBusIngressList = (params) => {
    return service({
        url: '/busIngress/getBusIngressList',
        method: 'get',
        params
    })
}

// @Summary 提交入库表
export const ingress = (data) => {
    return service({
        url: '/busIngress/ingress',
        method: 'post',
        data
    })
}