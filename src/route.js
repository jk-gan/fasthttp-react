
import React from 'react';
import { Route, Switch, BrowserRouter } from 'react-router-dom';
import { Provider } from 'mobx-react';
import NotFound from './views/errors/NotFound';
import './assets/js/GlobalStyle';

export default (
  <Provider>
    <BrowserRouter>
      <Switch>
        <Route path="/" component={NotFound} />
      </Switch>
    </BrowserRouter>
  </Provider>
);

