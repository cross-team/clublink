import React, { Component } from 'react';

import styles from './Button.module.scss';
import { Styling, withCSSModule } from './styling';

interface Props extends Styling {
  onClick?: () => void;
  className?: string;
}

export class Button extends Component<Props> {
  static defaultProps: Props = {
    styles: ['pink']
  };
  handleClick = () => {
    if (!this.props.onClick) {
      return;
    }
    this.props.onClick();
  };

  render() {
    return (
      <button
        className={`${withCSSModule(this.props.styles, styles)} ${
          styles.button
        } ${this.props.className}`}
        onClick={this.handleClick}
      >
        {this.props.children}
      </button>
    );
  }
}
