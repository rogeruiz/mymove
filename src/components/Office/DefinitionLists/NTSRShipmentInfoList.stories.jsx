import React from 'react';

import NTSRShipmentInfoList from './NTSRShipmentInfoList';

export default {
  title: 'Office Components/Shipment Info List',
  component: NTSRShipmentInfoList,
};

const info = {
  primeActualWeight: 2000,
  storageFacility: {
    address: {
      city: 'Anytown',
      country: 'USA',
      postalCode: '90210',
      state: 'OK',
      streetAddress1: '555 Main Ave',
      streetAddress2: 'Apartment 900',
    },
    facilityName: 'my storage',
    lotNumber: '2222',
  },
  serviceOrderNumber: '12341234',
  requestedDeliveryDate: '26 Mar 2020',
  destinationAddress: {
    streetAddress1: '441 SW Rio de la Plata Drive',
    city: 'Tacoma',
    state: 'WA',
    postalCode: '98421',
  },
  secondaryDeliveryAddress: {
    streetAddress1: '812 S 129th St',
    city: 'San Antonio',
    state: 'TX',
    postalCode: '78234',
  },
  agents: [
    {
      agentType: 'RECEIVING_AGENT',
      firstName: 'Kate',
      lastName: 'Smith',
      phone: '419-555-9999',
      email: 'ksmith@email.com',
    },
  ],
  counselorRemarks:
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam vulputate commodo erat. ' +
    'Morbi porta nibh nibh, ac malesuada tortor egestas.',
  customerRemarks: 'Ut enim ad minima veniam',
  tacType: 'HHG',
  sacType: 'NTS',
};

export const NTSRBasic = () => (
  <NTSRShipmentInfoList
    shipment={{
      counselorRemarks: info.counselorRemarks,
      requestedDeliveryDate: info.requestedDeliveryDate,
      storageFacility: info.storageFacility,
      destinationAddress: info.destinationAddress,
      tacType: info.tacType,
      sacType: info.sacType,
      primeActualWeight: info.primeActualWeight,
      serviceOrderNumber: info.serviceOrderNumber,
    }}
  />
);

export const NTSRMissingInfo = () => (
  <NTSRShipmentInfoList
    shipment={{
      requestedDeliveryDate: info.requestedDeliveryDate,
      destinationAddress: info.destinationAddress,
      tacType: info.tacType,
      sacType: info.sacType,
      primeActualWeight: info.primeActualWeight,
      serviceOrderNumber: info.serviceOrderNumber,
    }}
    errorIfMissing={['storageFacility']}
  />
);

export const WithAgents = () => (
  <NTSRShipmentInfoList
    shipment={{
      counselorRemarks: info.counselorRemarks,
      requestedDeliveryDate: info.requestedDeliveryDate,
      storageFacility: info.storageFacility,
      destinationAddress: info.destinationAddress,
      tacType: info.tacType,
      sacType: info.sacType,
      primeActualWeight: info.primeActualWeight,
      serviceOrderNumber: info.serviceOrderNumber,
      agents: [info.agents[0], info.agents[1]],
    }}
  />
);

export const WithRemarks = () => (
  <NTSRShipmentInfoList
    shipment={{
      requestedPickupDate: info.requestedPickupDate,
      pickupAddress: info.pickupAddress,
      destinationAddress: info.destinationAddress,
      counselorRemarks: info.counselorRemarks,
      customerRemarks: info.customerRemarks,
    }}
  />
);

export const WithAllInfo = () => (
  <NTSRShipmentInfoList
    shipment={{
      requestedDeliveryDate: info.requestedDeliveryDate,
      storageFacility: info.storageFacility,
      tacType: info.tacType,
      sacType: info.sacType,
      primeActualWeight: info.primeActualWeight,
      serviceOrderNumber: info.serviceOrderNumber,
      destinationAddress: info.destinationAddress,
      secondaryDeliveryAddress: info.secondaryDeliveryAddress,
      agents: [info.agents[0], info.agents[1]],
      counselorRemarks: info.counselorRemarks,
      customerRemarks: info.customerRemarks,
    }}
  />
);