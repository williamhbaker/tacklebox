import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import * as api from 'api';

// thunks

export const login = createAsyncThunk(
  'user/login',
  async (data, { rejectWithValue }) => {
    const result = await api.login(data);
    return result ? result : rejectWithValue();
  }
);

export const logout = createAsyncThunk(
  'user/logout',
  async (_, { rejectWithValue }) => {
    const result = await api.logout();
    return result ? result : rejectWithValue();
  }
);

// slice

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
      state.inProgress = false;
      state.initialized = true;
    },
    [logout.fulfilled]: (state, action) => {
      state.user = '';
      state.inProgress = false;
    },
  },
});

export const { receiveUser } = userSlice.actions;

export default userSlice.reducer;

// selectors

export const selectLoginInProgress = (state) => state.user.inProgress;
export const selectUser = (state) => state.user.user;
