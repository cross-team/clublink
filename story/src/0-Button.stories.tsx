import React from 'react';
import { Button } from '../../frontend/src/component/ui/Button';
import { action } from '@storybook/addon-actions';
import { text } from '@storybook/addon-knobs';

import styles from './0-Button.stories.module.scss';

export default {
  title: 'UI/Button',
  component: Button
};

export const pink = () => {
  return <Button onClick={action('click')}>{text('Label', 'Button')}</Button>;
};

export const blue = () => {
  return (
    <Button styles={['blue']} onClick={action('click')}>
      {text('Label', 'Button')}
    </Button>
  );
};

export const black = () => {
  return (
    <Button styles={['black']} onClick={action('click')}>
      {text('Label', 'Button')}
    </Button>
  );
};

export const red = () => {
  return (
    <Button styles={['red']} onClick={action('click')}>
      {text('Label', 'Button')}
    </Button>
  );
};

export const green = () => {
  return (
    <Button styles={['green']} onClick={action('click')}>
      {text('Label', 'Button')}
    </Button>
  );
};

export const fullWidth = () => {
  return (
    <Button styles={['pink', 'full-width']} onClick={action('click')}>
      <div className={styles.content}>{text('Label', 'Button')}</div>
    </Button>
  );
};

export const shadow = () => {
  return (
    <Button styles={['pink', 'shadow']} onClick={action('click')}>
      {text('Label', 'Button')}
    </Button>
  );
};
