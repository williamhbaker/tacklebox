import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import {
  BrowserRouter as Router,
  Route,
  Switch,
  Redirect,
} from 'react-router-dom';

import { checkStatus, selectInitialized } from 'features/user/userSlice';

import NavBar from 'components/navbar';
import LogInForm from 'components/login';
import LogOut from 'components/logout';
import Bins from 'components/bins';
import Loading from 'components/loading';

const App = () => {
  const dispatch = useDispatch();
  const initialized = useSelector(selectInitialized);

  useEffect(() => {
    dispatch(checkStatus());
  }, [dispatch]);

  return (
    <Router>
      {initialized ? (
        <>
          <NavBar />

          <Switch>
            <Route exact path="/login">
              <LogInForm />
            </Route>
            <Route exact path="/logout">
              <LogOut />
            </Route>
            <Route exact path="/bins">
              <Bins />
            </Route>
            <Route path="*">
              <Redirect to="/bins" />
            </Route>
          </Switch>
        </>
      ) : (
        <Loading />
      )}
    </Router>
  );
};

export default App;
