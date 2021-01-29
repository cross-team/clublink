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
        {this.urlData !== null && (
          <div className="short-link-usage-wrapper">
            <ShortLinkUsage
              shortLink={`${this.urlData.get('shortLink')}`}
              longLink={`${this.urlData.get('longLink')}`}
              qrCodeUrl={`${this.urlData
                .get('qrCodeURL')
                ?.replace(this.regex, '+')}`}
            />
          </div>
        )}
        <a
          href={`${this.urlData.get('shortLink')}`}
          className="heading"
          target="_blank"
        >
          <h1>
            🚀
            <span className="lightGreen">club</span>l
            <span className="lightGreen">.</span>
            ink/{this.urlData.get('alias')}
          </h1>
        </a>
        <p>Imagine a link impossible to remember:</p>
        <a href={`${this.urlData.get('longLink')}`} target="_blank">
          {this.urlData.get('longLink')}
        </a>
        <div className="buttons">
          <button
            onClick={() => {
              navigator.clipboard
                .writeText(`${this.urlData.get('shortLink')}`)
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
            copy club-link
          </button>
          <a className="button" href="/create">
            create new
          </a>
        </div>
      </div>
    );
  };
}