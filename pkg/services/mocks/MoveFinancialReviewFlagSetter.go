// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	appcontext "github.com/transcom/mymove/pkg/appcontext"

	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// MoveFinancialReviewFlagSetter is an autogenerated mock type for the MoveFinancialReviewFlagSetter type
type MoveFinancialReviewFlagSetter struct {
	mock.Mock
}

// SetFinancialReviewFlag provides a mock function with given fields: appCtx, moveID, eTag, remarks
func (_m *MoveFinancialReviewFlagSetter) SetFinancialReviewFlag(appCtx appcontext.AppContext, moveID uuid.UUID, eTag string, remarks string) (*models.Move, error) {
	ret := _m.Called(appCtx, moveID, eTag, remarks)

	var r0 *models.Move
	if rf, ok := ret.Get(0).(func(appcontext.AppContext, uuid.UUID, string, string) *models.Move); ok {
		r0 = rf(appCtx, moveID, eTag, remarks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Move)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(appcontext.AppContext, uuid.UUID, string, string) error); ok {
		r1 = rf(appCtx, moveID, eTag, remarks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}