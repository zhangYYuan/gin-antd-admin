import request from '@/utils/request';

export async function fakeAccountLogin(data) {
  return request('base/login', {
    method: 'POST',
    data,
  });
}
