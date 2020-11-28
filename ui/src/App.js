import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import NavBar from 'components/navbar';
import LogInForm from 'components/login';
import LogOut from 'components/logout';

function App() {
  return (
    <Router>
      <NavBar />

      <Switch>
        <Route exact path="/login">
          <LogInForm />
        </Route>
        <Route exact path="/logout">
          <LogOut />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
