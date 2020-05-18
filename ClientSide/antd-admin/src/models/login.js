import { stringify } from 'querystring';
import { history } from 'umi';
import { fakeAccountLogin } from '@/services/login';
import { setAuthority } from '@/utils/authority';
import { getPageQuery } from '@/utils/utils';
import { encryption } from '@/utils/enctypt'

const Model = {
  namespace: 'login',
  state: {
    status: undefined,
    currentUser: {},
  },
  effects: {
    *login({ payload }, { call, put }) {
      // 登录参数处理
      const { password } = payload
      const p = {
        ...payload,
        password: '123456',
      }
      const response = yield call(fakeAccountLogin, p);
      const data = response.data
      yield put({
        type: 'changeLoginStatus',
        payload: data,
      }); // Login successfully
      const token = data.token;
      localStorage.setItem('antd-pro-token', token);
      localStorage.setItem('antd-pro-user', JSON.stringify(data));
      if (token) {
        const urlParams = new URL(window.location.href);
        const params = getPageQuery();
        let { redirect } = params;

        if (redirect) {
          const redirectUrlParams = new URL(redirect);
          if (redirectUrlParams.origin === urlParams.origin) {
            redirect = redirect.substr(urlParams.origin.length);
            if (redirect.match(/^\/.*#/)) {
              redirect = redirect.substr(redirect.indexOf('#') + 1);
            }
          } else {
            window.location.href = '/';
            return;
          }
        }
        history.push(redirect || '/');
      }
    },

    logout() {
      const { redirect } = getPageQuery(); // Note: There may be security issues, please note

      if (window.location.pathname !== '/user/login' && !redirect) {
        history.replace({
          pathname: '/user/login',
          search: stringify({
            redirect: window.location.href,
          }),
        });
      }
    },
  },
  reducers: {
    changeLoginStatus(state, { payload }) {
      setAuthority(payload.currentAuthority);
      return { ...state, status: payload.data, type: payload.type };
    }
  },
};
export default Model;
