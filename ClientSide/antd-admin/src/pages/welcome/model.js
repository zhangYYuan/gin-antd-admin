import { queryTreatType } from './service'

const Model = {
  namespace: 'TreatModel',
  state: {
    list: []
  },
  effect: {
    * fetch({payload}, {call, put}) {
      const response = yield call(queryTreatType, payload);
      console.log(response)
      yield put({
        type: 'queryList'
      })
    }
  },
  reducers: {
    queryList (state, action){
      return {...state, list: action.payload}
    }
  }
}

export default Model
