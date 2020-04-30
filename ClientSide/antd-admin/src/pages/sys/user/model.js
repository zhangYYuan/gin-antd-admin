import { addFakeList } from '../service';


const Model = {
  namespace: 'listAndbasicList',
  state: {
    list: [],
  },
  effects: {
    * fetch({payload}, {call, put}) {
      console.log('------->fetch')
    },
  },
  reducers: {
    queryList(state, action) {
      return { ...state, list: action.payload };
    },
  }
}

export default Model;
