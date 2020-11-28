import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import * as api from 'api';

// thunks

// slice

const initialState = {
  activeBin: '',
  allBins: [],
};

// selectors

const binsSlice = createSlice({
  name: 'bins',
  initialState: [],
  reducers: {},
});

export default binsSlice;
