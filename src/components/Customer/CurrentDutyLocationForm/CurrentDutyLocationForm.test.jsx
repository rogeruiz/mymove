import React from 'react';
import { render, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

import CurrentDutyLocationForm from './CurrentDutyLocationForm';

describe('CurrentDutyLocationForm component', () => {
  it('renders the form input', async () => {
    const { getByLabelText } = render(
      <CurrentDutyLocationForm
        onSubmit={jest.fn()}
        onBack={jest.fn()}
        initialValues={{ current_location: {} }}
        newDutyLocation={{}}
      />,
    );
    await waitFor(() => {
      expect(getByLabelText('What is your current duty location?')).toBeInstanceOf(HTMLInputElement);
    });
  });

  it('keeps the next button disabled if the form is not filled out', async () => {
    const { getByRole } = render(
      <CurrentDutyLocationForm
        onSubmit={jest.fn()}
        onBack={jest.fn()}
        initialValues={{ current_location: {} }}
        newDutyLocation={{}}
      />,
    );
    const submitBtn = getByRole('button', { name: 'Next' });

    await waitFor(() => {
      expect(submitBtn).toHaveAttribute('disabled');
    });
  });

  it('shows an error message if trying to submit an invalid form', async () => {
    const onSubmit = jest.fn();
    const { getByRole, getAllByText } = render(
      <CurrentDutyLocationForm
        onSubmit={onSubmit}
        onBack={jest.fn()}
        initialValues={{
          current_location: {
            address: {
              city: 'Los Angeles',
              state: 'CA',
              postalCode: '90245',
            },
            name: 'Los Angeles AFB',
            id: 'testId',
          },
        }}
        newDutyLocation={{
          id: 'testId',
        }}
      />,
    );
    const submitBtn = getByRole('button', { name: 'Next' });

    await waitFor(() => {
      expect(submitBtn).toHaveAttribute('disabled');
      expect(
        getAllByText('You entered the same duty location for your origin and destination. Please change one of them.')
          .length,
      ).toBe(1);
    });
  });

  it('submits the form when its valid', async () => {
    const onSubmit = jest.fn();
    const { getByRole } = render(
      <CurrentDutyLocationForm
        onSubmit={onSubmit}
        onBack={jest.fn()}
        initialValues={{
          current_location: {
            address: {
              city: 'San Diego',
              state: 'CA',
              postalCode: '92104',
            },
            id: 'testId',
          },
        }}
        newDutyLocation={{ id: 'test' }}
      />,
    );
    const submitBtn = getByRole('button', { name: 'Next' });

    userEvent.click(submitBtn);

    await waitFor(() => {
      expect(onSubmit).toHaveBeenCalled();
    });
  });

  it('uses the onBack handler when the back button is clicked', async () => {
    const onBack = jest.fn();
    const { getByRole } = render(
      <CurrentDutyLocationForm
        onSubmit={jest.fn()}
        onBack={onBack}
        initialValues={{
          current_location: {},
        }}
      />,
    );
    const backBtn = getByRole('button', { name: 'Back' });

    userEvent.click(backBtn);

    await waitFor(() => {
      expect(onBack).toHaveBeenCalled();
    });
  });

  afterEach(jest.resetAllMocks);
});
