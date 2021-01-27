import React, { Component } from 'react';
import './PublishedPage.scss';

export class PublishedPage extends Component {
  private urlData = new URLSearchParams(window.location.search);

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
        <a href="/">or create a new club-link</a>
      </div>
    );
  };
}
