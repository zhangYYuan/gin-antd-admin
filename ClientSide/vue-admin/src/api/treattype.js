import service from '@/utils/request'

export const getList = (data) => {
    return service({
        url: "/treattype/getpagelist",
        method: 'get',
        params: data
    })
}


export const createTreatType = (data) => {
    return service({
        url: "/treattype/create",
        method: 'post',
        data
    })
}


export const getById = (data) => {
    return service({
        url: "/treattype/getById",
        method: 'get',
        data
    })
}




export const updateTreatType = (data) => {
    return service({
        url: "/treattype/update",
        method: 'put',
        data
    })
}




export const getAlls = (data) => {
    return service({
        url: "/treattype/getAllApis",
        method: 'post',
        data
    })
}

// @Tags Api
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Api true "删除api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
export const deleteTreatType = (data) => {
    return service({
        url: "/treattype/delete",
        method: 'delete',
        data
    })
}
