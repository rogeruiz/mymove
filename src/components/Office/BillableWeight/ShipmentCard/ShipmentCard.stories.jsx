import React from 'react';

import ShipmentCard from './ShipmentCard';

export default {
  title: 'Office Components/BillableWeightShipmentCard',
  component: ShipmentCard,
};

const props = {
  billableWeight: 4014,
  dateReweighRequested: new Date('1/1/2020').toISOString(),
  departedDate: new Date('12/25/2019').toISOString(),
  pickupAddress: {
    city: 'Rancho Santa Margarita',
    state: 'CA',
    postal_code: '92688',
  },
  destinationAddress: {
    city: 'West Springfield Town',
    state: 'MA',
    postal_code: '01089',
  },
  estimatedWeight: 5000,
  originalWeight: 4014,
  reweighRemarks: 'Unable to perform reweigh because shipment was already unloaded',
};

export const Card = () => (
  <div style={{ margin: '0 auto', height: '100%', width: 336 }}>
    <ShipmentCard {...props} />
  </div>
);