package serviceparamvaluelookups

import (
	"fmt"
	"math"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/unit"
)

// WeightBilledLookup does lookup on weight billed
type WeightBilledLookup struct {
	MTOShipment models.MTOShipment
}

func (r WeightBilledLookup) lookup(appCtx appcontext.AppContext, keyData *ServiceItemParamKeyData) (string, error) {
	var estimatedWeight *unit.Pound
	var originalWeight *unit.Pound

	switch keyData.MTOServiceItem.ReService.Code {
	case models.ReServiceCodeDOSHUT,
		models.ReServiceCodeDDSHUT,
		models.ReServiceCodeIOSHUT,
		models.ReServiceCodeIDSHUT:
		estimatedWeight = keyData.MTOServiceItem.EstimatedWeight

		if estimatedWeight == nil {
			// TODO: Do we need a different error -- is this a "normal" scenario?
			return "", fmt.Errorf("could not find estimated weight for MTOServiceItemID [%s]", keyData.MTOServiceItem.ID)
		}
		originalWeight = keyData.MTOServiceItem.ActualWeight

		if originalWeight == nil {
			// TODO: Do we need a different error -- is this a "normal" scenario?
			return "", fmt.Errorf("could not find actual weight for MTOServiceItemID [%s]", keyData.MTOServiceItem.ID)
		}
	default:
		// Make sure there's an estimated weight since that's nullable
		estimatedWeight = r.MTOShipment.PrimeEstimatedWeight

		// Make sure there's an actual weight since that's nullable
		originalWeight = r.MTOShipment.PrimeActualWeight
		if originalWeight == nil {
			// TODO: Do we need a different error -- is this a "normal" scenario?
			return "", fmt.Errorf("could not find actual weight for MTOShipmentID [%s]", r.MTOShipment.ID)
		}
	}

	var value string
	if estimatedWeight != nil {
		estimatedWeightCap := math.Round(float64(*estimatedWeight) * 1.10)
		if float64(*originalWeight) > estimatedWeightCap {
			value = applyMinimum(keyData.MTOServiceItem.ReService.Code, r.MTOShipment.ShipmentType, int(estimatedWeightCap))
		} else {
			value = applyMinimum(keyData.MTOServiceItem.ReService.Code, r.MTOShipment.ShipmentType, int(*originalWeight))
		}
	} else {
		value = applyMinimum(keyData.MTOServiceItem.ReService.Code, r.MTOShipment.ShipmentType, int(*originalWeight))
	}

	return value, nil
}

// Looks at code and applies minimum if necessary, otherwise returns weight passed in
func applyMinimum(code models.ReServiceCode, shipmentType models.MTOShipmentType, weight int) string {
	result := weight
	switch shipmentType {
	case models.MTOShipmentTypeInternationalUB:
		switch code {
		case models.ReServiceCodeIOSHUT,
			models.ReServiceCodeIDSHUT:
			if weight < 300 {
				result = 300
			}
		}
	default:
		switch code {
		case models.ReServiceCodeDLH,
			models.ReServiceCodeDSH,
			models.ReServiceCodeDOP,
			models.ReServiceCodeDDP,
			models.ReServiceCodeDOFSIT,
			models.ReServiceCodeDDFSIT,
			models.ReServiceCodeDOASIT,
			models.ReServiceCodeDDASIT,
			models.ReServiceCodeDOPSIT,
			models.ReServiceCodeDDDSIT,
			models.ReServiceCodeDPK,
			models.ReServiceCodeDUPK,
			models.ReServiceCodeDOSHUT,
			models.ReServiceCodeDDSHUT,
			models.ReServiceCodeIOOLH,
			models.ReServiceCodeICOLH,
			models.ReServiceCodeIOCLH,
			models.ReServiceCodeIHPK,
			models.ReServiceCodeIHUPK,
			models.ReServiceCodeIOFSIT,
			models.ReServiceCodeIDFSIT,
			models.ReServiceCodeIOASIT,
			models.ReServiceCodeIDASIT,
			models.ReServiceCodeIOPSIT,
			models.ReServiceCodeIDDSIT,
			models.ReServiceCodeIOSHUT,
			models.ReServiceCodeIDSHUT:
			if weight < 500 {
				result = 500
			}
		case models.ReServiceCodeIOOUB,
			models.ReServiceCodeICOUB,
			models.ReServiceCodeIOCUB,
			models.ReServiceCodeIUBPK,
			models.ReServiceCodeIUBUPK:
			if weight < 300 {
				result = 300
			}
		}
	}
	return fmt.Sprintf("%d", result)
}