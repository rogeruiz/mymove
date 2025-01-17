// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	appcontext "github.com/transcom/mymove/pkg/appcontext"

	notifications "github.com/transcom/mymove/pkg/notifications"

	testing "testing"
)

// NotificationSender is an autogenerated mock type for the NotificationSender type
type NotificationSender struct {
	mock.Mock
}

// SendNotification provides a mock function with given fields: appCtx, notification
func (_m *NotificationSender) SendNotification(appCtx appcontext.AppContext, notification notifications.Notification) error {
	ret := _m.Called(appCtx, notification)

	var r0 error
	if rf, ok := ret.Get(0).(func(appcontext.AppContext, notifications.Notification) error); ok {
		r0 = rf(appCtx, notification)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewNotificationSender creates a new instance of NotificationSender. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotificationSender(t testing.TB) *NotificationSender {
	mock := &NotificationSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
