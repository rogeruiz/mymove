// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"
)

// MoveRouter is an autogenerated mock type for the MoveRouter type
type MoveRouter struct {
	mock.Mock
}

// Approve provides a mock function with given fields: move
func (_m *MoveRouter) Approve(move *models.Move) error {
	ret := _m.Called(move)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Move) error); ok {
		r0 = rf(move)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cancel provides a mock function with given fields: reason, move
func (_m *MoveRouter) Cancel(reason string, move *models.Move) error {
	ret := _m.Called(reason, move)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *models.Move) error); ok {
		r0 = rf(reason, move)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteServiceCounseling provides a mock function with given fields: move
func (_m *MoveRouter) CompleteServiceCounseling(move *models.Move) error {
	ret := _m.Called(move)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Move) error); ok {
		r0 = rf(move)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendToOfficeUserToReviewNewServiceItems provides a mock function with given fields: move
func (_m *MoveRouter) SendToOfficeUserToReviewNewServiceItems(move *models.Move) error {
	ret := _m.Called(move)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Move) error); ok {
		r0 = rf(move)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Submit provides a mock function with given fields: move
func (_m *MoveRouter) Submit(move *models.Move) error {
	ret := _m.Called(move)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Move) error); ok {
		r0 = rf(move)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}