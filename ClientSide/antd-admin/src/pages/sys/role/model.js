// import { formatterTreeNode } from '@/utils';
import { formatterTreeNode } from '@/utils/index';
import { queryDepart } from '../service';

// deptCode: "11000000"
// deptCodeIds: null
// deptLeader: null
// deptLeave: 1
// deptName: "北京华夏名鉴科技有限公司"

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
      const data = formatterTreeNode(action.payload, 'deptName', 'deptCode')
      return { ...state, departList: data };
    },
  }
}

export default Model;

