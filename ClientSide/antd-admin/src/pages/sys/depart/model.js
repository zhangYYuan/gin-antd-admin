import utils from '@/utils/index';
import { queryDepart } from '../service';

const Model = {
  namespace: 'sysDepart',
  state: {
    departList: [],
  },
  effects: {
    * fetch({payload}, {call, put}) {
      const response = yield call(queryDepart, payload);
      if (response.resultCode === 200) {
        yield put({
          type: 'queryDepart',
          payload: response.resultBody,
        });
      }
    },
  },
  reducers: {
    queryDepart(state, action) {
      const data = utils.formatterTreeNode(action.payload, 'deptName', 'deptCode')
      return { ...state, departList: data };
    },
  }
}

export default Model;

