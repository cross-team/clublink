import React, { ChangeEvent, Component, createRef } from 'react';
import './TextField.scss';

interface IProps {
  disabled?: boolean;
  className?: string;
  text?: string;
  placeHolder?: string;
  onChange?: (text: string) => void;
  onBlur?: () => void;
  onFocus?: () => void;
}

export class TextField extends Component<IProps, any> {
  textInput = createRef<HTMLInputElement>();

  render = () => {
    return (
      <input
        ref={this.textInput}
        className={`text-field ${this.props.className && this.props.className}`}
        type={'text'}
        value={this.props.text}
        onChange={this.handleChange}
        onBlur={this.handleBlur}
        onFocus={this.handleFocus}
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

  focus = () => {
    this.textInput.current!.focus();
  };
}