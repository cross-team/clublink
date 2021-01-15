import React from 'react';
import { render } from '@testing-library/react';
import { DEFAULT_DURATION, Toast } from './Toast';

describe('Toast component', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  afterEach(() => {
    jest.clearAllTimers();
  });

  test('should render without crash', () => {
    render(<Toast />);
  });

  test('should show content correctly when triggered to show', () => {
    const toastRef = React.createRef<Toast>();
    const toastMessage = 'Toast Message';
    const { container } = render(<Toast ref={toastRef} />);

    expect(container.textContent).not.toContain(toastMessage);
    toastRef.current!.notify(toastMessage, 1000);
    expect(container.textContent).toContain(toastMessage);
  });

  test('should show content for default duration when no delay given', () => {
    const toastRef = React.createRef<Toast>();
    const toastMessage = 'Toast Message';
    const { container } = render(<Toast ref={toastRef} />);

    expect(container.textContent).not.toContain(toastMessage);
    toastRef.current!.notify(toastMessage);
    expect(container.textContent).toContain(toastMessage);

    jest.advanceTimersByTime(DEFAULT_DURATION - 1);
    expect(container.textContent).toContain(toastMessage);

    jest.advanceTimersByTime(1);
    expect(container.textContent).not.toContain(toastMessage);
  });

  test('should automatically hide content after delay', () => {
    const toastRef = React.createRef<Toast>();
    const toastMessage = 'Toast Message';
    const { container } = render(<Toast ref={toastRef} />);
    const TOTAL_DURATION = 2000;
    const HALF_TIME = 1000;

    expect(container.textContent).not.toContain(toastMessage);
    toastRef.current!.notify(toastMessage, TOTAL_DURATION);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).toContain(toastMessage);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).not.toContain(toastMessage);
  });

  test('second notify call should replace first toast', () => {
    const toastRef = React.createRef<Toast>();
    const firstToastMessage = 'First Toast Message';
    const secondToastMessage = 'Second Toast Message';
    const { container } = render(<Toast ref={toastRef} />);
    const TOTAL_DURATION = 2000;
    const HALF_TIME = 1000;

    expect(container.textContent).not.toContain(firstToastMessage);
    toastRef.current!.notify(firstToastMessage, TOTAL_DURATION);

    jest.advanceTimersByTime(HALF_TIME);
    // second notify before the first one closes
    toastRef.current!.notify(secondToastMessage, TOTAL_DURATION);
    expect(container.textContent).not.toContain(firstToastMessage);
    expect(container.textContent).toContain(secondToastMessage);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).toContain(secondToastMessage);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).not.toContain(secondToastMessage);
  });
});
