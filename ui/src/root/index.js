import React from 'react';
import { Provider } from 'react-redux';
import {
  BrowserRouter as Router,
  Route,
  Switch,
  Redirect,
} from 'react-router-dom';

import store from './store';

import NavBar from 'components/navbar';

const Root = () => (
  <Provider store={store}>
    <Router>
      <NavBar />

      <Switch></Switch>
    </Router>
  </Provider>
);

export default Root;
