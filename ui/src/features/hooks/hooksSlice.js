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

export const deleteHook = createAsyncThunk(
  'hooks/delete',
  async (hookID, { rejectWithValue }) => {
    const result = await api.deleteHook(hookID);
    return result ? hookID : rejectWithValue();
  }
);

// slice

const initialState = {
  inProgress: false,
  hooks: [],
  hookDetails: [],
  activeHook: '',
};

const hooksSlice = createSlice({
  name: 'hooks',
  initialState: initialState,
  reducers: {
    activateHook(state, action) {
      state.activeHook = action.payload;
    },
  },
  extraReducers: {
    [activateBin.pending]: (state, action) => {
      state.inProgress = true;
    },
    [activateBin.fulfilled]: (state, action) => {
      state.hooks = action.payload.map((h) => h.ID);
      state.hookDetails = action.payload.map((h) => ({
        id: h.ID,
        content: h.Content,
        created: h.Created,
      }));
      state.inProgress = false;
    },
    [activateBin.rejected]: (state, action) => {
      state.inProgress = false;
    },
    [deleteHook.pending]: (state, action) => {
      state.inProgress = true;
    },
    [deleteHook.fulfilled]: (state, action) => {
      state.hooks = state.hooks.filter((h) => h !== action.payload);
      state.hookDetails = state.hookDetails.filter(
        (h) => h.id !== action.payload
      );
      state.activeHook = '';
      state.inProgress = false;
    },
    [deleteHook.rejected]: (state, action) => {
      state.inProgress = false;
    },
  },
});

export const { activateHook } = hooksSlice.actions;

export default hooksSlice.reducer;

// selectors

export const selectHooks = (state) => state.hooks.hooks;
export const selectHooksInProgress = (state) => state.hooks.inProgress;

export const selectActiveHook = (state) => {
  return state.hooks.activeHook;
};

export const selectActiveHookDetails = (state) => {
  return state.hooks.hookDetails.filter(
    (d) => d.id === state.hooks.activeHook
  )[0];
};
