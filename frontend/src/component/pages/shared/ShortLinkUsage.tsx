import React, { Component } from 'react';

import './ShortLinkUsage.scss';

interface Props {
  shortLink: string;
  longLink: string;
  qrCodeUrl: string;
}

export class ShortLinkUsage extends Component<Props> {
  render() {
    return (
      <div className={'short-link-usage'}>
        {/* <div>
          You can now paste&nbsp;
          <a target={'_blank'} href={this.props.shortLink}>
            {this.props.shortLink}
          </a>
          &nbsp;in your browser to visit
          <br />
          {this.props.longLink}.
        </div> */}
        <div className={'qr-code'}>
          <img alt="Q R code for link" src={this.props.qrCodeUrl} />
        </div>
      </div>
    );
  }
}
