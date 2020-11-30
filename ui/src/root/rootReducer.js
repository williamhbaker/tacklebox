import { combineReducers } from 'redux';

import { logout } from 'features/user/userSlice';

import userReducer from 'features/user/userSlice';
import binsReducer from 'features/bins/binsSlice';
import hooksReducer from 'features/hooks/hooksSlice';

const appReducer = combineReducers({
  user: userReducer,
  bins: binsReducer,
  hooks: hooksReducer,
});

const rootReducer = (state, action) => {
  if (action.type === logout.fulfilled().type) {
    state = undefined;
  }

  return appReducer(state, action);
};

export default rootReducer;
