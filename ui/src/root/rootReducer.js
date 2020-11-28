import { combineReducers } from 'redux';

import userReducer from 'features/user/userSlice';
import binsReducer from 'features/bins/binsSlice';
import hooksReducer from 'features/hooks/hooksSlice';

export default combineReducers({
  user: userReducer,
  bins: binsReducer,
  hooks: hooksReducer,
});
