import { queryMenu } from '../service';
import utils  from '@/utils';

const Model = {
  namespace: 'sysMenu',
  state: {
    menuList: [],
  },
  effects: {
    * fetch({payload}, {call, put}) {
      const response = yield call(queryMenu, payload);
      if (response.resultCode === 200) {
        yield put({
          type: 'queryMenu',
          payload: response.resultBody,
        });
      }
    },
  },
  reducers: {
    queryMenu(state, action) {
      const data = utils.formatterTreeSelect(action.payload, 'menuName', 'menuCode')
      return { ...state, menuList: data };
    },
  }
}

export default Model;

