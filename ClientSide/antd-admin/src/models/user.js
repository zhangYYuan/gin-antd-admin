import { query as queryUsers } from '@/services/user';

const UserModel = {
  namespace: 'user',
  state: {
    currentUser: {},
  },
  effects: {
    *fetch(_, { call, put }) {
      const response = yield call(queryUsers);
      yield put({
        type: 'save',
        payload: response,
      });
    },

    *fetchCurrent(_, { put }) {
      const user = JSON.parse(localStorage.getItem('antd-pro-user'));
      console.log('fetchCurrent----->');
      yield put({
        type: 'saveCurrentUser',
        payload: user,
      });
    },
  },
  reducers: {
    saveCurrentUser(state, action) {
      console.log('saveCurrentUser----->');
      return { ...state, currentUser: action.payload || {} };
    },

    changeNotifyCount(
      state = {
        currentUser: {},
      },
      action,
    ) {
      return {
        ...state,
        currentUser: {
          ...state.currentUser,
          notifyCount: action.payload.totalCount,
          unreadCount: action.payload.unreadCount,
        },
      };
    },
  },
};
export default UserModel;
