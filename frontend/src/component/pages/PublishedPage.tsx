import React, { Component } from 'react';
import './PublishedPage.scss';
import { ShortLinkUsage } from './shared/ShortLinkUsage';

export class PublishedPage extends Component {
  private urlData: URLSearchParams = new URLSearchParams(
    window.location.search
  );
  private regex = / /g;

  componentDidMount() {
    console.log(`${this.urlData.get('qrCodeURL')?.replace(this.regex, '+')}`);
  }

  render = () => {
    return (
      <div className="published">
        <h1>
          🚀
          <span className="lightGreen">club</span>l
          <span className="lightGreen">.</span>
          ink/{this.urlData.get('alias')}
        </h1>
        <button
          onClick={() => {
            navigator.clipboard
              .writeText(`clubl.ink/${this.urlData.get('alias')}`)
              .then(
                function() {
                  /* clipboard successfully set */
                },
                function() {
                  /* clipboard write failed */
                }
              );
          }}
        >
          copy
        </button>
        {this.urlData !== null && (
          <div className={'short-link-usage-wrapper'}>
            <ShortLinkUsage
              shortLink={`${this.urlData.get('shortLink')}`}
              longLink={`${this.urlData.get('longLink')}`}
              qrCodeUrl={`${this.urlData
                .get('qrCodeURL')
                ?.replace(this.regex, '+')}`}
            />
          </div>
        )}
        <a href="/">or create a new club-link</a>
      </div>
    );
  };
}
