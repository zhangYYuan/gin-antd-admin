import request from '@/utils/request';

export async function addMenu(params) {
  return request('/api/menu/addAntMenu', {
    method: 'POST',
    data: params
  });
}

export async function removeRule(params) {
  return request('/api/rule', {
    method: 'POST',
    data: { ...params, method: 'delete' },
  });
}
export async function addRule(params) {
  return request('sys/role/page?pageNum=2&pageSize=10', {
    method: 'POST',
    data: { ...params, method: 'post' },
  });
}
export async function queryRole(params) {
  return request('/sys/role/page', {
    method: 'get',
    params
  });
}
