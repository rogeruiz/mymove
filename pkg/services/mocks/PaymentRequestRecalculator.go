// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	appcontext "github.com/transcom/mymove/pkg/appcontext"

	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// PaymentRequestRecalculator is an autogenerated mock type for the PaymentRequestRecalculator type
type PaymentRequestRecalculator struct {
	mock.Mock
}

// RecalculatePaymentRequest provides a mock function with given fields: appCtx, paymentRequestID
func (_m *PaymentRequestRecalculator) RecalculatePaymentRequest(appCtx appcontext.AppContext, paymentRequestID uuid.UUID) (*models.PaymentRequest, error) {
	ret := _m.Called(appCtx, paymentRequestID)

	var r0 *models.PaymentRequest
	if rf, ok := ret.Get(0).(func(appcontext.AppContext, uuid.UUID) *models.PaymentRequest); ok {
		r0 = rf(appCtx, paymentRequestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PaymentRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(appcontext.AppContext, uuid.UUID) error); ok {
		r1 = rf(appCtx, paymentRequestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}