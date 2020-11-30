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

export const signUp = createAsyncThunk(
  'user/signUp',
  async (data, { rejectWithValue }) => {
    const result = await api.signUp(data);
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

export const checkStatus = createAsyncThunk(
  'user/checkStatus',
  async (_, { rejectWithValue }) => {
    const result = await api.checkStatus();
    return result ? result : rejectWithValue();
  }
);

// slice

const initialState = {
  user: '',
  inProgress: false,
  initialized: false,
  message: '',
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setMessage(state, action) {
      state.message = action.payload;
    },
  },
  extraReducers: {
    [login.pending]: (state, action) => {
      state.user = '';
      state.inProgress = true;
    },
    [login.fulfilled]: (state, action) => {
      state.user = action.payload.message;
      state.inProgress = false;
      state.message = '';
    },
    [login.rejected]: (state, action) => {
      state.user = '';
      state.inProgress = false;
      state.message = '';
    },
    [signUp.pending]: (state, action) => {
      state.inProgress = true;
    },
    [signUp.fulfilled]: (state, action) => {
      state.inProgress = false;
    },
    [signUp.rejected]: (state, action) => {
      state.inProgress = false;
    },
    [logout.fulfilled]: (state, action) => {
      state.user = '';
      state.inProgress = false;
    },
    [checkStatus.pending]: (state, action) => {
      state.inProgress = true;
    },
    [checkStatus.fulfilled]: (state, action) => {
      state.user = action.payload.message;
      state.initialized = true;
      state.inProgress = false;
    },
    [checkStatus.rejected]: (state, action) => {
      state.user = '';
      state.initialized = true;
      state.inProgress = false;
    },
  },
});

export const { setMessage } = userSlice.actions;

export default userSlice.reducer;

// selectors

export const selectLoginInProgress = (state) => state.user.inProgress;
export const selectInitialized = (state) => state.user.initialized;
export const selectUser = (state) => state.user.user;
export const selectMessage = (state) => state.user.message;
