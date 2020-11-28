import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import * as api from 'api';

// thunks

export const getHooks = createAsyncThunk(
  'hooks/getHooks',
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
    [getHooks.pending]: (state, action) => {
      state.inProgress = true;
    },
    [getHooks.fulfilled]: (state, action) => {
      state.hooks = action.payload;
      state.inProgress = false;
    },
    [getHooks.rejected]: (state, action) => {
      state.inProgress = false;
    },
  },
});

export default hooksSlice.reducer;

// selectors

export const selectHooks = (state) => state.hooks.hooks;

// export const selectBinsLoadingInProgress = (state) => state.bins.inProgress;
// export const selectBins = (state) => state.bins.bins;
