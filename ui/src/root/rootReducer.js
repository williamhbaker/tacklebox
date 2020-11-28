import { combineReducers } from 'redux';

import userReducer from 'features/user/userSlice';

export default combineReducers({
  user: userReducer,
});
