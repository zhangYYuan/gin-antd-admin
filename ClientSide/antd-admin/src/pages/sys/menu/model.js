import { queryMenu } from '../service';
import {connect} from "dva";

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
      console.log('------->fetch', action.payload)
      return { ...state, menuList: action.payload };
    },
  }
}

export default Model;

