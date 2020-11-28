import { combineReducers } from 'redux';

import binsReducer from 'features/bins/binsSlice';

export default combineReducers({
  bins: binsReducer,
});
