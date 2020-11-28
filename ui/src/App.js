import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import NavBar from 'components/navbar';
import LogInForm from 'components/login';

function App() {
  return (
    <Router>
      <NavBar />

      <Switch>
        <Route exact path="/login">
          <LogInForm />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
