import service from '@/utils/request'

export const getList = (data) => {
    console.log('getList--data: ' + JSON.stringify(data))
    return service({
        url: "/treattype/getpagelist",
        method: 'post',
        data
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
        url: "/treattype/getbyid",
        method: 'post',
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
        url: "/treattype/getalls",
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

export function statusFilter(status) {
    const statusMap = {
        1: 'success',
        0: 'danger'
    }
    return statusMap[status]
}

export function statusNameFilter(status) {
    const statusMap = {
        1: '正常',
        0: '已删除'
    }
    return statusMap[status]
}

export function levelFilter(level) {
    const levelsMap = {
        5: 'success',
        6: 'danger',
        1: 'dark',
        2: 'warning',
        3: 'plain',
        4: 'info',
    }
    return levelsMap[level]
}

export function levelNameFilter(level) {
    const levelsMap = {
        1: '一阶',
        2: '二阶',
        3: '三阶'
    }
    return levelsMap[level]
}

export function regionFilter(region) {
    const regionsMap = {
        3: 'success',
        4: 'danger',
        5: 'dark',
        6: 'warning',
        2: 'plain',
        1: 'info',
    }
    return regionsMap[region]
}

export function regionNameFilter(region) {
    const regionsMap = {
        1: '国内',
        2: '海外',
    }
    return regionsMap[region]
}
