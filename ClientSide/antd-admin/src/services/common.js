import request from '@/utils/request';

export async function queryProvince(params) {
  return request('/wms/area/getProvinceById', {
    params,
  });
}

export async function queryCity(params) {
  return request('/wms/area/getCityById', {
    params,
  });
}

export async function queryCountry(params) {
  return request('/wms/area/getCountryById', {
    params
  })
}


