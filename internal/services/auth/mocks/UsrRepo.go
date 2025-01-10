// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/EtoNeAnanasbI95/auth-grpc-demo/internal/domain/models"
	mock "github.com/stretchr/testify/mock"
)

// UsrRepo is an autogenerated mock type for the UsrRepo type
type UsrRepo struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: ctx, email, passwordHash
func (_m *UsrRepo) AddUser(ctx context.Context, email string, passwordHash []byte) (int64, error) {
	ret := _m.Called(ctx, email, passwordHash)

	if len(ret) == 0 {
		panic("no return value specified for AddUser")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) (int64, error)); ok {
		return rf(ctx, email, passwordHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) int64); ok {
		r0 = rf(ctx, email, passwordHash)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []byte) error); ok {
		r1 = rf(ctx, email, passwordHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, email
func (_m *UsrRepo) GetUser(ctx context.Context, email string) (*models.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserWithId provides a mock function with given fields: ctx, uid
func (_m *UsrRepo) GetUserWithId(ctx context.Context, uid int64) (*models.User, error) {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for GetUserWithId")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.User, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.User); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsrRepo creates a new instance of UsrRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsrRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsrRepo {
	mock := &UsrRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
