/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';

import { Home } from './index';

import { MockProviders } from 'testUtils';

export default {
  title: 'Customer Components / Pages / Home',
};

const uploadOrdersProps = {
  serviceMember: {
    id: 'testServiceMemberId',
    first_name: 'John',
    last_name: 'Lee',
    current_location: {
      name: 'Fort Knox',
      transportation_office: {
        name: 'Test Transportation Office Name',
        phone_lines: ['555-555-5555'],
      },
    },
    weight_allotment: {},
  },
  showLoggedInUser() {},
  loadMTOShipments() {},
  history: { push: () => {}, goBack: () => {} },
  getSignedCertification() {},
  updateShipmentList() {},
  mtoShipments: [],
  mtoShipment: {},
  isLoggedIn: true,
  loggedInUserIsLoading: false,
  loggedInUserSuccess: true,
  isProfileComplete: true,
  orders: {},
  location: {},
  move: {
    locator: 'XYZ890',
    status: 'DRAFT',
  },
  uploadedOrderDocuments: [],
  uploadedAmendedOrderDocuments: [],
};

const shipmentSelectionProps = {
  ...uploadOrdersProps,
  serviceMember: {
    ...uploadOrdersProps.serviceMember,
    weight_allotment: {
      total_weight_self: 10000,
    },
  },
  orders: {
    ...uploadOrdersProps.orders,
    new_duty_location: {
      name: 'NAS Jacksonville',
    },
    origin_duty_location: {
      name: 'NAS Norfolk',
    },
    report_by_date: '2020-12-25',
  },
  uploadedOrderDocuments: [
    {
      id: 'file1',
      filename: 'Uploaded_Orders.pdf',
    },
    {
      id: 'file2',
      filename: 'Supporting_Documentation_Screenshot.png',
    },
  ],
};

const withShipmentProps = {
  ...shipmentSelectionProps,
  mtoShipments: [
    {
      id: 'testShipment1',
      shipmentType: 'HHG',
      createdAt: '2020-12-24',
    },
    {
      id: 'testMove',
      shipmentType: 'PPM',
    },
    {
      id: 'testShipment2',
      shipmentType: 'PPM',
      createdAt: '24 December 2020',
      ppmShipment: {
        id: 'ppmShipmentIncomplete',
        hasRequestedAdvance: null,
      },
    },
    {
      id: 'completePPM',
      shipmentType: 'PPM',
      createdAt: '24 December 2020',
      ppmShipment: {
        id: 'ppmShipmentComplete',
        hasRequestedAdvance: true,
      },
    },
  ],
};

const submittedProps = {
  ...withShipmentProps,
  move: {
    ...withShipmentProps.move,
    status: 'SUBMITTED',
    submitted_at: '2020-12-24',
  },
};

const amendedOrderProps = {
  ...submittedProps,
  move: {
    ...submittedProps.move,
    status: 'APPROVALS REQUESTED',
  },
  uploadedAmendedOrderDocuments: [
    {
      id: 'file3',
      filename: 'Amended_Orders.pdf',
    },
  ],
};

export const Step2 = () => {
  return (
    <MockProviders>
      <div className="grid-container usa-prose">
        <Home {...uploadOrdersProps} />
      </div>
    </MockProviders>
  );
};

export const Step3 = () => {
  return (
    <MockProviders>
      <div className="grid-container usa-prose">
        <Home {...shipmentSelectionProps} />
      </div>
    </MockProviders>
  );
};

export const Step4 = () => {
  return (
    <MockProviders>
      <div className="grid-container usa-prose">
        <Home {...withShipmentProps} />
      </div>
    </MockProviders>
  );
};

export const SubmittedMoveWithoutPPM = () => {
  let { mtoShipments } = submittedProps;
  mtoShipments = mtoShipments.slice(0, 1);
  const props = { ...submittedProps, mtoShipments };
  return (
    <MockProviders>
      <div className="grid-container usa-prose">
        <Home {...props} />
      </div>
    </MockProviders>
  );
};

export const SubmittedMoveWithPPM = () => {
  return (
    <MockProviders>
      <div className="grid-container usa-prose">
        <Home {...submittedProps} />
      </div>
    </MockProviders>
  );
};

export const AmendedOrders = () => {
  return (
    <MockProviders>
      <div className="grid-container usa-prose">
        <Home {...amendedOrderProps} />
      </div>
    </MockProviders>
  );
};
