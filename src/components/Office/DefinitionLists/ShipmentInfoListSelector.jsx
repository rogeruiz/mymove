import React from 'react';
import * as PropTypes from 'prop-types';

import { ShipmentShape } from 'types/shipment';
import ShipmentInfoList from 'components/Office/DefinitionLists/ShipmentInfoList';
import NTSRShipmentInfoList from 'components/Office/DefinitionLists/NTSRShipmentInfoList';
import { SHIPMENT_OPTIONS } from 'shared/constants';

const ShipmentInfoListSelector = ({
  className,
  shipment,
  isExpanded,
  warnIfMissing,
  errorIfMissing,
  showWhenCollapsed,
  shipmentType,
}) => {
  switch (shipmentType) {
    case SHIPMENT_OPTIONS.HHG:
      return <ShipmentInfoList className={className} shipment={shipment} shipmentType={shipmentType} />;
    case SHIPMENT_OPTIONS.NTSR:
      return (
        <NTSRShipmentInfoList
          className={className}
          shipment={shipment}
          isExpanded={isExpanded}
          warnIfMissing={warnIfMissing}
          errorIfMissing={errorIfMissing}
          showWhenCollapsed={showWhenCollapsed}
        />
      );
    default:
      return (
        <ShipmentInfoList
          className={className}
          shipment={shipment}
          shipmentType={shipmentType}
          isExpanded={isExpanded}
        />
      );
  }
};

ShipmentInfoListSelector.propTypes = {
  className: PropTypes.string,
  shipment: ShipmentShape.isRequired,
  isExpanded: PropTypes.bool,
  warnIfMissing: PropTypes.arrayOf(PropTypes.string),
  errorIfMissing: PropTypes.arrayOf(PropTypes.string),
  showWhenCollapsed: PropTypes.arrayOf(PropTypes.string),
  shipmentType: PropTypes.oneOf([
    SHIPMENT_OPTIONS.HHG,
    SHIPMENT_OPTIONS.HHG_SHORTHAUL_DOMESTIC,
    SHIPMENT_OPTIONS.HHG_LONGHAUL_DOMESTIC,
    SHIPMENT_OPTIONS.NTS,
    SHIPMENT_OPTIONS.NTSR,
  ]),
};

ShipmentInfoListSelector.defaultProps = {
  shipmentType: SHIPMENT_OPTIONS.HHG,
  className: '',
  isExpanded: false,
  warnIfMissing: [],
  errorIfMissing: [],
  showWhenCollapsed: [],
};

export default ShipmentInfoListSelector;