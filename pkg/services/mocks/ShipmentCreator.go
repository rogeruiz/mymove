// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	appcontext "github.com/transcom/mymove/pkg/appcontext"

	models "github.com/transcom/mymove/pkg/models"

	testing "testing"
)

// ShipmentCreator is an autogenerated mock type for the ShipmentCreator type
type ShipmentCreator struct {
	mock.Mock
}

// CreateShipment provides a mock function with given fields: appCtx, shipment
func (_m *ShipmentCreator) CreateShipment(appCtx appcontext.AppContext, shipment *models.MTOShipment) (*models.MTOShipment, error) {
	ret := _m.Called(appCtx, shipment)

	var r0 *models.MTOShipment
	if rf, ok := ret.Get(0).(func(appcontext.AppContext, *models.MTOShipment) *models.MTOShipment); ok {
		r0 = rf(appCtx, shipment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOShipment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(appcontext.AppContext, *models.MTOShipment) error); ok {
		r1 = rf(appCtx, shipment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewShipmentCreator creates a new instance of ShipmentCreator. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewShipmentCreator(t testing.TB) *ShipmentCreator {
	mock := &ShipmentCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}