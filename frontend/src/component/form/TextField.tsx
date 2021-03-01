import React, { ChangeEvent, Component, createRef } from 'react';
import './TextField.scss';

interface IProps {
  disabled?: boolean;
  id?: string;
  className?: string;
  text?: string;
  placeHolder?: string;
  aria?: string;
  describedBy?: string;
  onChange?: (text: string) => void;
  onBlur?: () => void;
  onFocus?: () => void;
  onKeyPress?: (e: any) => void;
}

export class TextField extends Component<IProps, any> {
  textInput = createRef<HTMLInputElement>();

  render = () => {
    return (
      <input
        aria-label={this.props.aria}
        aria-describedby={this.props.describedBy}
        ref={this.textInput}
        id={this.props.id}
        className={`text-field ${this.props.className && this.props.className}`}
        type={'text'}
        value={this.props.text}
        onChange={this.handleChange}
        onBlur={this.handleBlur}
        onFocus={this.handleFocus}
        onKeyPress={this.handleKeyPress}
        placeholder={this.props.placeHolder}
        disabled={this.props.disabled}
      />
    );
  };

  handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const text = event.target.value;
    if (!this.props.onChange) {
      return;
    }
    this.props.onChange(text);
  };

  handleBlur = () => {
    if (!this.props.onBlur) {
      return;
    }
    this.props.onBlur();
  };

  handleFocus = () => {
    if (!this.props.onFocus) {
      return;
    }
    this.props.onFocus();
  };

  handleKeyPress = (e: any) => {
    if (!this.props.onKeyPress) {
      return;
    }
    this.props.onKeyPress(e);
  };

  focus = () => {
    this.textInput.current!.focus();
  };
}
