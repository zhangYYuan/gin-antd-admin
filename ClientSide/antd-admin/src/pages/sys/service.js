import request from '@/utils/request';

/**
 * 查询角色列表
 * @param params
 * @returns {Promise<any>}
 */
export async function queryRole(params) {
  return request('/sys/role/page', {
    method: 'get',
    params
  });
}


/**
 * 查询菜单
 * @returns {Promise<void>}
 */
export async function queryMenu(params) {
  return request('sys/menu/list', {
    method: 'get',
    params
  });
}

export async function addMenu(data) {
  return request('menu/add', {
    method: 'post',
    data
  });
}

/**
 * 部门管理
 * @param params
 * @returns {Promise<any>}
 */
export async function queryDepart(params) {
  return request('sys/dept/listSysDeptList', {
    method: 'get',
    params
  });
}
