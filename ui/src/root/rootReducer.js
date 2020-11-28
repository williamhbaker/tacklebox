import { combineReducers } from 'redux';

import userReducer from 'features/user/userSlice';
import binsReducer from 'features/bins/binsSlice';

export default combineReducers({
  user: userReducer,
  bins: binsReducer,
});
