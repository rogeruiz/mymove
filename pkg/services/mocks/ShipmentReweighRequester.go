// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// ShipmentReweighRequester is an autogenerated mock type for the ShipmentReweighRequester type
type ShipmentReweighRequester struct {
	mock.Mock
}

// RequestShipmentReweigh provides a mock function with given fields: ctx, shipmentID
func (_m *ShipmentReweighRequester) RequestShipmentReweigh(ctx context.Context, shipmentID uuid.UUID) (*models.Reweigh, error) {
	ret := _m.Called(ctx, shipmentID)

	var r0 *models.Reweigh
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *models.Reweigh); ok {
		r0 = rf(ctx, shipmentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Reweigh)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, shipmentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}