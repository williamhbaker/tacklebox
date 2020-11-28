import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import * as api from 'api';

// thunks

export const getBins = createAsyncThunk(
  'bins/getBins',
  async (_, { rejectWithValue }) => {
    const result = await api.getBins();
    return result ? result : rejectWithValue();
  }
);

export const createBin = createAsyncThunk(
  'bins/createBin',
  async (_, { dispatch, rejectWithValue }) => {
    const result = await api.createBin();
    if (!result) {
      rejectWithValue();
    } else {
      dispatch(getBins());
    }
  }
);

export const destroyBin = createAsyncThunk(
  'bins/destroyBin',
  async (binID, { dispatch, rejectWithValue }) => {
    const result = await api.destroyBin(binID);
    if (!result) {
      rejectWithValue();
    } else {
      dispatch(getBins());
    }
  }
);

// slice

const initialState = {
  inProgress: false,
  bins: [],
};

const binsSlice = createSlice({
  name: 'bins',
  initialState: initialState,
  reducers: {},
  extraReducers: {
    [getBins.pending]: (state, action) => {
      state.inProgress = true;
    },
    [getBins.fulfilled]: (state, action) => {
      state.bins = action.payload;
      state.inProgress = false;
    },
    [getBins.rejected]: (state, action) => {
      state.inProgress = false;
    },
    [createBin.pending]: (state, action) => {
      state.inProgress = true;
    },
    [createBin.rejected]: (state, action) => {
      state.inProgress = false;
    },
  },
});

export default binsSlice.reducer;

// selectors

export const selectBinsLoadingInProgress = (state) => state.bins.inProgress;
export const selectBins = (state) => state.bins.bins;
