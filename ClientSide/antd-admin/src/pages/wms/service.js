import request from '@/utils/request';

export async function queryStoreList(params) {
  return request('/wms/store/pageList', {
    params
  });
}
