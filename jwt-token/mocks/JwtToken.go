// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"

	model "github.com/itbasis/go-jwt-auth/model"

	time "time"
)

// JwtToken is an autogenerated mock type for the JwtToken type
type JwtToken struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: _a0, _a1
func (_m *JwtToken) CreateAccessToken(_a0 context.Context, _a1 model.SessionUser) (string, *time.Time, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 *time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SessionUser) (string, *time.Time, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.SessionUser) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.SessionUser) *time.Time); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*time.Time)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, model.SessionUser) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateRefreshToken provides a mock function with given fields: _a0, _a1
func (_m *JwtToken) CreateRefreshToken(_a0 context.Context, _a1 model.SessionUser) (string, *time.Time, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 *time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SessionUser) (string, *time.Time, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.SessionUser) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.SessionUser) *time.Time); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*time.Time)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, model.SessionUser) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateTokenCustomDuration provides a mock function with given fields: _a0, _a1, _a2
func (_m *JwtToken) CreateTokenCustomDuration(_a0 context.Context, _a1 model.SessionUser, _a2 time.Duration) (string, *time.Time, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	var r1 *time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SessionUser, time.Duration) (string, *time.Time, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.SessionUser, time.Duration) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.SessionUser, time.Duration) *time.Time); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*time.Time)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, model.SessionUser, time.Duration) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Parse provides a mock function with given fields: ctx, tokenString
func (_m *JwtToken) Parse(ctx context.Context, tokenString string) (*model.SessionUser, error) {
	ret := _m.Called(ctx, tokenString)

	var r0 *model.SessionUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.SessionUser, error)); ok {
		return rf(ctx, tokenString)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.SessionUser); ok {
		r0 = rf(ctx, tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SessionUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSecretKey provides a mock function with given fields: secretKey, signMethod
func (_m *JwtToken) SetSecretKey(secretKey []byte, signMethod jwt.SigningMethod) {
	_m.Called(secretKey, signMethod)
}

type mockConstructorTestingTNewJwtToken interface {
	mock.TestingT
	Cleanup(func())
}

// NewJwtToken creates a new instance of JwtToken. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJwtToken(t mockConstructorTestingTNewJwtToken) *JwtToken {
	mock := &JwtToken{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
