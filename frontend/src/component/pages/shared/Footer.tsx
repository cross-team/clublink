import React, { Component } from 'react';

import './Footer.scss';
import { UIFactory } from '../../UIFactory';

interface Props {
  uiFactory: UIFactory;
  authorName: string;
  authorPortfolio: string;
  version: string;
  onShowChangeLogBtnClick: () => void;
}

export class Footer extends Component<Props> {
  render() {
    return (
      <footer>
        <div className={'center'}>
          <div className={'row'}>Loving clubl.ink? ❤️ Love us back!</div>
          <div className={'row app-version'}>
            <a>@mpaiva</a> <a>@paiva</a> <a>@sebastian</a> <a>@laurabries</a>
            <a>@jomonti</a>
          </div>
          {/* <div className={'row'}>
            {this.props.uiFactory.createViewChangeLogButton({
              onClick: this.props.onShowChangeLogBtnClick
            })}
          </div> */}
        </div>
      </footer>
    );
  }
}
