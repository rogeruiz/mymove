import React from 'react';
import { render, screen, waitFor, within } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { generatePath } from 'react-router';

import { primeSimulatorRoutes } from '../../../constants/routes';
import { MockProviders } from '../../../testUtils';
import { usePrimeSimulatorGetMove } from '../../../hooks/queries';
import { updatePrimeMTOShipment } from '../../../services/primeApi';

import PrimeUIShipmentUpdate from './PrimeUIShipmentUpdate';

const shipmentId = 'ce01a5b8-9b44-4511-8a8d-edb60f2a4aee';
const moveId = '9c7b255c-2981-4bf8-839f-61c7458e2b4d';

const mockPush = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useParams: jest.fn().mockReturnValue({
    moveCode: 'LR4T8V',
    moveCodeOrID: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
    shipmentId: 'ce01a5b8-9b44-4511-8a8d-edb60f2a4aee',
  }),
  useHistory: () => ({
    push: mockPush,
  }),
}));

jest.mock('services/primeApi', () => ({
  ...jest.requireActual('services/primeApi'),
  updatePrimeMTOShipment: jest.fn().mockImplementation(() => Promise.resolve()),
}));

jest.mock('hooks/queries', () => ({
  usePrimeSimulatorGetMove: jest.fn(),
}));
const approvedMoveTaskOrder = {
  moveTaskOrder: {
    availableToPrimeAt: '2021-10-18T18:24:41.235Z',
    createdAt: '2021-10-18T18:24:41.362Z',
    eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4zNjIxNjRa',
    excessWeightAcknowledgedAt: null,
    excessWeightQualifiedAt: null,
    excessWeightUploadId: null,
    id: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
    moveCode: 'LR4T8V',
    mtoShipments: [
      {
        actualPickupDate: '2020-03-16',
        agents: [
          {
            agentType: 'RELEASING_AGENT',
            createdAt: '2021-10-18T18:24:41.521Z',
            eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS41MjE4NzNa',
            email: 'test@test.email.com',
            firstName: 'Test',
            id: 'f2619e1b-7729-4b97-845d-6ae1ebe299f2',
            lastName: 'Agent',
            mtoShipmentID: 'ce01a5b8-9b44-4511-8a8d-edb60f2a4aee',
            phone: '202-555-9301',
            updatedAt: '2021-10-18T18:24:41.521Z',
          },
        ],
        approvedDate: '2021-10-18',
        createdAt: '2021-10-18T18:24:41.377Z',
        customerRemarks: 'Please treat gently',
        destinationAddress: {
          city: 'Fairfield',
          country: 'US',
          eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4zNzI3NDJa',
          id: 'bfe61147-5fd7-426e-b473-54ccf77bde35',
          postalCode: '94535',
          state: 'CA',
          streetAddress1: '987 Any Avenue',
          streetAddress2: 'P.O. Box 9876',
          streetAddress3: 'c/o Some Person',
        },
        eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4zNzc5Nzha',
        firstAvailableDeliveryDate: null,
        id: 'ce01a5b8-9b44-4511-8a8d-edb60f2a4aee',
        moveTaskOrderID: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
        pickupAddress: {
          city: 'Beverly Hills',
          country: 'US',
          eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4zNjc3Mjda',
          id: 'cf159eca-162c-4131-84a0-795e684416a6',
          postalCode: '90210',
          state: 'CA',
          streetAddress1: '123 Any Street',
          streetAddress2: 'P.O. Box 12345',
          streetAddress3: 'c/o Some Person',
        },
        primeActualWeight: 2000,
        primeEstimatedWeight: 1400,
        primeEstimatedWeightRecordedDate: null,
        requestedPickupDate: '2020-03-15',
        requiredDeliveryDate: null,
        scheduledPickupDate: '2020-03-16',
        secondaryDeliveryAddress: {
          city: null,
          postalCode: null,
          state: null,
          streetAddress1: null,
        },
        secondaryPickupAddress: {
          city: null,
          postalCode: null,
          state: null,
          streetAddress1: null,
        },
        shipmentType: 'HHG_LONGHAUL_DOMESTIC',
        status: 'APPROVED',
        updatedAt: '2021-10-18T18:24:41.377Z',
        mtoServiceItems: null,
      },
    ],
    order: {
      customer: {
        branch: 'AIR_FORCE',
        dodID: '5917531070',
        eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4xNDIxNTZa',
        email: 'leo_spaceman_sm@example.com',
        firstName: 'Leo',
        id: 'e2de409b-edb9-42af-b50f-564458e08ada',
        lastName: 'Spacemen',
        phone: '555-555-5555',
        userID: 'ae204f8a-6222-45a1-9b79-e2d52441b4f2',
      },
      customerID: 'e2de409b-edb9-42af-b50f-564458e08ada',
      destinationDutyStation: {
        address: {
          city: 'Augusta',
          country: 'United States',
          eTag: 'MjAyMS0xMC0xOFQxODoyMzoxMi4zMTQzNDZa',
          id: '5ac95be8-0230-47ea-90b4-b0f6f60de364',
          postalCode: '30813',
          state: 'GA',
          streetAddress1: 'Fort Gordon',
        },
        addressID: '5ac95be8-0230-47ea-90b4-b0f6f60de364',
        id: '2d5ada83-e09a-47f8-8de6-83ec51694a86',
        name: 'Fort Gordon',
      },
      eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4yMzAxMVo=',
      entitlement: {
        authorizedWeight: 8000,
        dependentsAuthorized: true,
        eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4xNzc0MjZa',
        id: '46ee60c2-9b17-44c7-9202-15a84327fc2f',
        nonTemporaryStorage: true,
        organizationalClothingAndIndividualEquipment: true,
        privatelyOwnedVehicle: true,
        proGearWeight: 2000,
        proGearWeightSpouse: 500,
        requiredMedicalEquipmentWeight: 1000,
        storageInTransit: 2,
        totalDependents: 1,
        totalWeight: 5000,
      },
      id: '8cda4825-283c-4910-89f4-1741e2fd9cb7',
      linesOfAccounting: 'F8E1',
      orderNumber: 'ORDER3',
      originDutyStation: {
        address: {
          city: 'Des Moines',
          country: 'US',
          eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS4yMDgyNjha',
          id: 'dbbee525-9c88-40c1-a549-6330b35972d2',
          postalCode: '50309',
          state: 'IA',
          streetAddress1: '987 Other Avenue',
          streetAddress2: 'P.O. Box 1234',
          streetAddress3: 'c/o Another Person',
        },
        addressID: 'dbbee525-9c88-40c1-a549-6330b35972d2',
        id: '0ecd8fb1-0551-44c8-a15e-83c5f4e3ae0f',
        name: 'XOXhgDSIRS',
      },
      rank: 'E_1',
      reportByDate: '2018-08-01',
    },
    orderID: '8cda4825-283c-4910-89f4-1741e2fd9cb7',
    paymentRequests: [
      {
        eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS41Nzc2OTha',
        id: '532ec513-8297-44b3-91a8-5167650b2869',
        isFinal: false,
        moveTaskOrderID: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
        paymentRequestNumber: '3301-9920-1',
        paymentServiceItems: [
          {
            eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS42Mzc5MzJa',
            id: '8fdf0b3a-c102-4084-84fe-22903f20470b',
            mtoServiceItemID: '8829fb28-69c1-45d7-98bc-c724478d5106',
            paymentRequestID: '532ec513-8297-44b3-91a8-5167650b2869',
            referenceID: '3301-9920-8fdf0b3a',
            status: 'REQUESTED',
          },
        ],
        status: 'PENDING',
      },
    ],
    ppmType: 'PARTIAL',
    referenceId: '3301-9920',
    updatedAt: '2021-10-18T18:24:41.362Z',
    mtoServiceItems: [
      {
        reServiceCode: 'STEST',
        eTag: 'MjAyMS0xMC0xOFQxODoyNDo0MS41MzE0NjRa',
        id: '8829fb28-69c1-45d7-98bc-c724478d5106',
        modelType: 'MTOServiceItemBasic',
        moveTaskOrderID: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
        mtoShipmentID: 'ce01a5b8-9b44-4511-8a8d-edb60f2a4aee',
        reServiceName: 'Test Service',
        status: 'APPROVED',
      },
    ],
  },
};

/*
const tooApprovedMoveDetailsQuery = {
  ...newMoveDetailsQuery,
  moveTaskOrder: {
    id: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
    ordersId: '1',
    status: MOVE_STATUSES.APPROVED,
  },
};

 */

/*
   mtoShipments: ...mtoShipments,
    mtoShipments: [
      {
        id: 'ce01a5b8-9b44-4511-8a8d-edb60f2a4aee',
        moveTaskOrderID: '9c7b255c-2981-4bf8-839f-61c7458e2b4d',
        status: shipmentStatuses.APPROVED,
      }
 */

const updateShipmentURL = generatePath(primeSimulatorRoutes.UPDATE_SHIPMENT_PATH, {
  moveCodeOrID: moveId,
  shipmentId,
});
const moveDetailsURL = generatePath(primeSimulatorRoutes.VIEW_MOVE_PATH, { moveCodeOrID: moveId });

const mockedComponent = (
  <MockProviders initialEntries={[updateShipmentURL]}>
    <PrimeUIShipmentUpdate />
  </MockProviders>
);

const readyReturnValue = {
  ...approvedMoveTaskOrder,
  isLoading: false,
  isError: false,
  isSuccess: true,
};

const loadingReturnValue = {
  ...approvedMoveTaskOrder,
  isLoading: true,
  isError: false,
  isSuccess: false,
};

const errorReturnValue = {
  ...approvedMoveTaskOrder,
  isLoading: false,
  isError: true,
  isSuccess: false,
};

describe('Update Shipment Page', () => {
  it('renders the page without errors', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(readyReturnValue);

    render(mockedComponent);

    await waitFor(() => expect(screen.getByText('Shipment Dates')).toBeInTheDocument());
    await waitFor(() => expect(screen.getByText('Shipment Weights')).toBeInTheDocument());
    await waitFor(() => expect(screen.getByText('Shipment Addresses')).toBeInTheDocument());
  });

  it('renders the Loading Placeholder when the query is still loading', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(loadingReturnValue);

    render(mockedComponent);

    const h2 = await screen.getByRole('heading', { name: 'Loading, please wait...', level: 2 });
    expect(h2).toBeInTheDocument();
  });

  it('renders the Something Went Wrong component when the query errors', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(errorReturnValue);

    render(mockedComponent);

    const errorMessage = await screen.getByText(/Something went wrong./);
    expect(errorMessage).toBeInTheDocument();
  });

  it('navigates the user to the home page when the cancel button is clicked', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(readyReturnValue);

    render(<PrimeUIShipmentUpdate />);

    const cancel = screen.getByRole('button', { name: 'Cancel' });
    userEvent.click(cancel);

    await waitFor(() => {
      expect(mockPush).toHaveBeenCalledWith(moveDetailsURL);
    });
  });
});

describe('Displays the shipment information to update', () => {
  it('displays the shipment information', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(readyReturnValue);

    render(<PrimeUIShipmentUpdate />);

    const shipmentDatesHeader = screen.getByRole('heading', { name: 'Shipment Dates', level: 2 });
    expect(shipmentDatesHeader).toBeInTheDocument();
    const updateShipmentContainer = shipmentDatesHeader.parentElement;

    await waitFor(() => {
      expect(
        within(updateShipmentContainer).getByRole('heading', {
          name: 'Shipment Weights',
          level: 2,
        }),
      ).toBeInTheDocument();
      expect(
        within(updateShipmentContainer).getByRole('heading', {
          name: 'Shipment Addresses',
          level: 2,
        }),
      ).toBeInTheDocument();
    });
  });
  /*
it('displays the submit button disabled', async () => {

usePrimeSimulatorGetMove.mockReturnValue(missingPrimeUpdates);

render(<PrimeUIShipmentUpdate />);

expect(screen.getByRole('button', { name: 'Save' })).toBeDisabled();
expect(
  screen.getByText(
    'At least one basic service item or shipment service item is required to create a payment request',
  ),
).toBeInTheDocument();

  });
   */
  it('displays the submit button active', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(readyReturnValue);

    render(<PrimeUIShipmentUpdate />);

    await waitFor(() => {
      expect(screen.getByRole('button', { name: 'Save' })).toBeEnabled();
      expect(screen.getByText(/123 Any Street/)).toBeInTheDocument();
    });
  });
});

describe('successful submission of form', () => {
  it('calls history router back to move details', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(readyReturnValue);
    updatePrimeMTOShipment.mockReturnValue({});

    render(<PrimeUIShipmentUpdate />);

    const actualPickupDateInput = await screen.findByLabelText('Actual pickup');
    userEvent.type(actualPickupDateInput, '2021-10-20');

    const actualWeightInput = screen.getByLabelText(/Actual weight/);
    userEvent.type(actualWeightInput, '10000');

    const saveButton = await screen.getByRole('button', { name: 'Save' });

    expect(saveButton).not.toBeDisabled();
    userEvent.click(saveButton);

    await waitFor(() => {
      expect(mockPush).toHaveBeenCalledWith(moveDetailsURL);
    });
  });

  /*
  it('update shipment', async () => {
    usePrimeSimulatorGetMove.mockReturnValue(missingPrimeUpdates);
    updatePrimeMTOShipment.mockReturnValue({});

    render(<PrimeUIShipmentUpdate />);

    const actualPickupDateInput = await screen.findByLabelText('Actual pickup');
    userEvent.type(actualPickupDateInput, '2021-10-20');

    const actualWeightInput = screen.getByLabelText(/Actual weight/);
    userEvent.type(actualWeightInput, "10000")

    //const saveButton = await expect(screen.getByRole('button', { name: 'Save' })).toBeEnabled();
    const saveButton = await screen.getByRole('button', { name: 'Save' });

    expect(saveButton).not.toBeDisabled();
    userEvent.click(saveButton);

    await waitFor(() => {
      expect(mockPush).toHaveBeenCalledWith(moveDetailsURL);
    });
  });
   */
});