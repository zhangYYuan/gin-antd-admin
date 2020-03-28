import request from '@/utils/request';

export async function fakeAccountLogin(params) {
  return request('/auth/oauth/token', {
    method: 'POST',
    params,
  });
}
export async function getFakeCaptcha(mobile) {
  return request(`/api/login/captcha?mobile=${mobile}`);
}
