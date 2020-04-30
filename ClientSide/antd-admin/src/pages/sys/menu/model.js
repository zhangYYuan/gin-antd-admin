import { addFakeList } from '../service';


const Model = {
  namespace: 'sysMenu',
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
