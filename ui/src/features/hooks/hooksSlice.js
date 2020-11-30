import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import * as api from 'api';

// thunks

export const activateBin = createAsyncThunk(
  'hooks/activateBin',
  async (binID, { rejectWithValue }) => {
    const result = await api.getHooks(binID);
    return result ? result : rejectWithValue();
  }
);

// slice

const initialState = {
  inProgress: false,
  hooks: [],
};

const hooksSlice = createSlice({
  name: 'hooks',
  initialState: initialState,
  reducers: {},
  extraReducers: {
    [activateBin.pending]: (state, action) => {
      state.inProgress = true;
    },
    [activateBin.fulfilled]: (state, action) => {
      state.hooks = action.payload;
      state.inProgress = false;
    },
    [activateBin.rejected]: (state, action) => {
      state.inProgress = false;
    },
  },
});

export default hooksSlice.reducer;

// selectors

export const selectHooks = (state) => state.hooks.hooks;
export const selectHooksInProgress = (state) => state.hooks.inProgress;
