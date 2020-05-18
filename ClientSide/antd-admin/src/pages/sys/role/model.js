import utils  from '@/utils'
import { queryRole } from "../service";

const Model = {
  namespace: 'sysRole',
  state: {
    roleList: [],
  },
  effects: {
    * fetch({payload}, {call, put}) {
      const response = yield call(queryRole, payload);
      if (response.resultCode === 200) {
        yield put({
          type: 'queryRole',
          payload: response.resultBody.data,
        });
      }
    },
  },
  reducers: {
    queryRole(state, action) {
      const data = utils.formatterSelect(action.payload, 'roleName')
      console.log(data)
      return { ...state, roleList: data };
    },
  }
}

export default Model;

