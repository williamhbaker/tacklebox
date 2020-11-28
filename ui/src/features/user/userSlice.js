import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import * as api from 'api';

export const login = createAsyncThunk(
  'user/login',
  async (data, { rejectWithValue }) => {
    const result = await api.login(data);
    return result ? result : rejectWithValue();
  }
);

let initialState = {
  user: '',
  inProgress: false,
  initialized: false,
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {},
  extraReducers: {
    [login.pending]: (state, action) => {
      state.user = '';
      state.inProgress = true;
    },
    [login.fulfilled]: (state, action) => {
      state.user = action.payload.message;
      state.inProgress = false;
      state.initialized = true;
    },
    [login.rejected]: (state, action) => {
      state.user = '';
      state.inProgress = true;
      state.initialized = true;
    },
  },
});

export const { receiveUser } = userSlice.actions;

export default userSlice.reducer;
