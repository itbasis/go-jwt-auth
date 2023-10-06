// Code generated by mockery v2.35.2. DO NOT EDIT.

package jwttoken

import (
	context "context"

	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"

	model "github.com/itbasis/go-jwt-auth/v2/model"

	time "time"
)

// MockJwtToken is an autogenerated mock type for the JwtToken type
type MockJwtToken struct {
	mock.Mock
}

type MockJwtToken_Expecter struct {
	mock *mock.Mock
}

func (_m *MockJwtToken) EXPECT() *MockJwtToken_Expecter {
	return &MockJwtToken_Expecter{mock: &_m.Mock}
}

// CreateAccessToken provides a mock function with given fields: _a0, _a1
func (_m *MockJwtToken) CreateAccessToken(_a0 context.Context, _a1 model.SessionUser) (string, *time.Time, error) {
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

// MockJwtToken_CreateAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccessToken'
type MockJwtToken_CreateAccessToken_Call struct {
	*mock.Call
}

// CreateAccessToken is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.SessionUser
func (_e *MockJwtToken_Expecter) CreateAccessToken(_a0 interface{}, _a1 interface{}) *MockJwtToken_CreateAccessToken_Call {
	return &MockJwtToken_CreateAccessToken_Call{Call: _e.mock.On("CreateAccessToken", _a0, _a1)}
}

func (_c *MockJwtToken_CreateAccessToken_Call) Run(run func(_a0 context.Context, _a1 model.SessionUser)) *MockJwtToken_CreateAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SessionUser))
	})
	return _c
}

func (_c *MockJwtToken_CreateAccessToken_Call) Return(_a0 string, _a1 *time.Time, _a2 error) *MockJwtToken_CreateAccessToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockJwtToken_CreateAccessToken_Call) RunAndReturn(run func(context.Context, model.SessionUser) (string, *time.Time, error)) *MockJwtToken_CreateAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRefreshToken provides a mock function with given fields: _a0, _a1
func (_m *MockJwtToken) CreateRefreshToken(_a0 context.Context, _a1 model.SessionUser) (string, *time.Time, error) {
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

// MockJwtToken_CreateRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRefreshToken'
type MockJwtToken_CreateRefreshToken_Call struct {
	*mock.Call
}

// CreateRefreshToken is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.SessionUser
func (_e *MockJwtToken_Expecter) CreateRefreshToken(_a0 interface{}, _a1 interface{}) *MockJwtToken_CreateRefreshToken_Call {
	return &MockJwtToken_CreateRefreshToken_Call{Call: _e.mock.On("CreateRefreshToken", _a0, _a1)}
}

func (_c *MockJwtToken_CreateRefreshToken_Call) Run(run func(_a0 context.Context, _a1 model.SessionUser)) *MockJwtToken_CreateRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SessionUser))
	})
	return _c
}

func (_c *MockJwtToken_CreateRefreshToken_Call) Return(_a0 string, _a1 *time.Time, _a2 error) *MockJwtToken_CreateRefreshToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockJwtToken_CreateRefreshToken_Call) RunAndReturn(run func(context.Context, model.SessionUser) (string, *time.Time, error)) *MockJwtToken_CreateRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTokenCustomDuration provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockJwtToken) CreateTokenCustomDuration(_a0 context.Context, _a1 model.SessionUser, _a2 time.Duration) (string, *time.Time, error) {
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

// MockJwtToken_CreateTokenCustomDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTokenCustomDuration'
type MockJwtToken_CreateTokenCustomDuration_Call struct {
	*mock.Call
}

// CreateTokenCustomDuration is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.SessionUser
//   - _a2 time.Duration
func (_e *MockJwtToken_Expecter) CreateTokenCustomDuration(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockJwtToken_CreateTokenCustomDuration_Call {
	return &MockJwtToken_CreateTokenCustomDuration_Call{Call: _e.mock.On("CreateTokenCustomDuration", _a0, _a1, _a2)}
}

func (_c *MockJwtToken_CreateTokenCustomDuration_Call) Run(run func(_a0 context.Context, _a1 model.SessionUser, _a2 time.Duration)) *MockJwtToken_CreateTokenCustomDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SessionUser), args[2].(time.Duration))
	})
	return _c
}

func (_c *MockJwtToken_CreateTokenCustomDuration_Call) Return(_a0 string, _a1 *time.Time, _a2 error) *MockJwtToken_CreateTokenCustomDuration_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockJwtToken_CreateTokenCustomDuration_Call) RunAndReturn(run func(context.Context, model.SessionUser, time.Duration) (string, *time.Time, error)) *MockJwtToken_CreateTokenCustomDuration_Call {
	_c.Call.Return(run)
	return _c
}

// Parse provides a mock function with given fields: ctx, tokenString
func (_m *MockJwtToken) Parse(ctx context.Context, tokenString string) (*model.SessionUser, error) {
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

// MockJwtToken_Parse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parse'
type MockJwtToken_Parse_Call struct {
	*mock.Call
}

// Parse is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenString string
func (_e *MockJwtToken_Expecter) Parse(ctx interface{}, tokenString interface{}) *MockJwtToken_Parse_Call {
	return &MockJwtToken_Parse_Call{Call: _e.mock.On("Parse", ctx, tokenString)}
}

func (_c *MockJwtToken_Parse_Call) Run(run func(ctx context.Context, tokenString string)) *MockJwtToken_Parse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockJwtToken_Parse_Call) Return(_a0 *model.SessionUser, _a1 error) *MockJwtToken_Parse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockJwtToken_Parse_Call) RunAndReturn(run func(context.Context, string) (*model.SessionUser, error)) *MockJwtToken_Parse_Call {
	_c.Call.Return(run)
	return _c
}

// SetSecretKey provides a mock function with given fields: secretKey, signMethod
func (_m *MockJwtToken) SetSecretKey(secretKey []byte, signMethod jwt.SigningMethod) {
	_m.Called(secretKey, signMethod)
}

// MockJwtToken_SetSecretKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSecretKey'
type MockJwtToken_SetSecretKey_Call struct {
	*mock.Call
}

// SetSecretKey is a helper method to define mock.On call
//   - secretKey []byte
//   - signMethod jwt.SigningMethod
func (_e *MockJwtToken_Expecter) SetSecretKey(secretKey interface{}, signMethod interface{}) *MockJwtToken_SetSecretKey_Call {
	return &MockJwtToken_SetSecretKey_Call{Call: _e.mock.On("SetSecretKey", secretKey, signMethod)}
}

func (_c *MockJwtToken_SetSecretKey_Call) Run(run func(secretKey []byte, signMethod jwt.SigningMethod)) *MockJwtToken_SetSecretKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(jwt.SigningMethod))
	})
	return _c
}

func (_c *MockJwtToken_SetSecretKey_Call) Return() *MockJwtToken_SetSecretKey_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockJwtToken_SetSecretKey_Call) RunAndReturn(run func([]byte, jwt.SigningMethod)) *MockJwtToken_SetSecretKey_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockJwtToken creates a new instance of MockJwtToken. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockJwtToken(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockJwtToken {
	mock := &MockJwtToken{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
