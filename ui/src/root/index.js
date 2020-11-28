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
import LogInForm from 'components/login';

const Root = () => (
  <Provider store={store}>
    <Router>
      <NavBar />

      <Switch>
        <Route exact path="/login">
          <LogInForm />
        </Route>
      </Switch>
    </Router>
  </Provider>
);

export default Root;
