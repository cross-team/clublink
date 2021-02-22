import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import { ShortLinkService } from '../service/ShortLink.service';
import { UIFactory } from './UIFactory';
import { NotFoundPage } from './pages/NotFoundPage';

interface IProps {
  shortLinkService: ShortLinkService;
  uiFactory: UIFactory;
}

export class App extends Component<IProps> {
  componentDidMount() {
    let nunito = document.createElement('style');
    nunito.append(
      `@import url('https://fonts.googleapis.com/css2?family=Nunito:ital,wght@0,200;0,300;0,400;0,600;0,700;0,800;0,900;1,200;1,300;1,400;1,600;1,700;1,800;1,900&display=swap');`
    );
    document.head.appendChild(nunito);
  }

  render = () => {
    return (
      <Router>
        <Switch>
          <Route
            path={'/a/enter-code'}
            exact
            render={({ location, history }) =>
              this.props.uiFactory.createHomePage(location, history, 'visit')
            }
          />
          <Route
            path={'/a/create'}
            exact
            render={({ location, history }) =>
              this.props.uiFactory.createHomePage(location, history, 'create')
            }
          />
          <Route
            path={'/a/favorites'}
            exact
            render={({ location, history }) =>
              this.props.uiFactory.createHomePage(
                location,
                history,
                'favorites'
              )
            }
          />
          <Route
            path={'/a/published'}
            exact
            render={() => {
              return this.props.uiFactory.createPublishedPage();
            }}
          />
          <Route
            path={'/a/admin'}
            exact
            render={() => {
              return this.props.uiFactory.createAdminPage();
            }}
          />
          <Route
            path={'/:alias'}
            render={({ match }) => {
              let alias = match.params['alias'];
              window.location.href = this.props.shortLinkService.aliasToBackendLink(
                alias
              );
              return <div />;
            }}
          />
          <Route
            path={'/'}
            exact
            render={({ location, history }) => {
              window.location.assign('/a/enter-code');
              return this.props.uiFactory.createHomePage(
                location,
                history,
                'visit'
              );
            }}
          />
          <Route component={NotFoundPage} />
        </Switch>
      </Router>
    );
  };
}
