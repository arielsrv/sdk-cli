// Code generated by mockery. DO NOT EDIT.

package services

import (
	model "github.com/arielsrv/sdk-cli/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// MockTemplateService is an autogenerated mock type for the TemplateService type
type MockTemplateService struct {
	mock.Mock
}

type MockTemplateService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTemplateService) EXPECT() *MockTemplateService_Expecter {
	return &MockTemplateService_Expecter{mock: &_m.Mock}
}

// CreateTemplate provides a mock function with given fields: templateName, appName
func (_m *MockTemplateService) CreateTemplate(templateName string, appName string) error {
	ret := _m.Called(templateName, appName)

	if len(ret) == 0 {
		panic("no return value specified for CreateTemplate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(templateName, appName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTemplateService_CreateTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTemplate'
type MockTemplateService_CreateTemplate_Call struct {
	*mock.Call
}

// CreateTemplate is a helper method to define mock.On call
//   - templateName string
//   - appName string
func (_e *MockTemplateService_Expecter) CreateTemplate(templateName interface{}, appName interface{}) *MockTemplateService_CreateTemplate_Call {
	return &MockTemplateService_CreateTemplate_Call{Call: _e.mock.On("CreateTemplate", templateName, appName)}
}

func (_c *MockTemplateService_CreateTemplate_Call) Run(run func(templateName string, appName string)) *MockTemplateService_CreateTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockTemplateService_CreateTemplate_Call) Return(_a0 error) *MockTemplateService_CreateTemplate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTemplateService_CreateTemplate_Call) RunAndReturn(run func(string, string) error) *MockTemplateService_CreateTemplate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAvailableLanguages provides a mock function with given fields:
func (_m *MockTemplateService) GetAvailableLanguages() []model.Language {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAvailableLanguages")
	}

	var r0 []model.Language
	if rf, ok := ret.Get(0).(func() []model.Language); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Language)
		}
	}

	return r0
}

// MockTemplateService_GetAvailableLanguages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAvailableLanguages'
type MockTemplateService_GetAvailableLanguages_Call struct {
	*mock.Call
}

// GetAvailableLanguages is a helper method to define mock.On call
func (_e *MockTemplateService_Expecter) GetAvailableLanguages() *MockTemplateService_GetAvailableLanguages_Call {
	return &MockTemplateService_GetAvailableLanguages_Call{Call: _e.mock.On("GetAvailableLanguages")}
}

func (_c *MockTemplateService_GetAvailableLanguages_Call) Run(run func()) *MockTemplateService_GetAvailableLanguages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTemplateService_GetAvailableLanguages_Call) Return(_a0 []model.Language) *MockTemplateService_GetAvailableLanguages_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTemplateService_GetAvailableLanguages_Call) RunAndReturn(run func() []model.Language) *MockTemplateService_GetAvailableLanguages_Call {
	_c.Call.Return(run)
	return _c
}

// GetTemplate provides a mock function with given fields: name
func (_m *MockTemplateService) GetTemplate(name string) (*model.Template, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetTemplate")
	}

	var r0 *model.Template
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Template, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Template); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Template)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTemplateService_GetTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplate'
type MockTemplateService_GetTemplate_Call struct {
	*mock.Call
}

// GetTemplate is a helper method to define mock.On call
//   - name string
func (_e *MockTemplateService_Expecter) GetTemplate(name interface{}) *MockTemplateService_GetTemplate_Call {
	return &MockTemplateService_GetTemplate_Call{Call: _e.mock.On("GetTemplate", name)}
}

func (_c *MockTemplateService_GetTemplate_Call) Run(run func(name string)) *MockTemplateService_GetTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTemplateService_GetTemplate_Call) Return(_a0 *model.Template, _a1 error) *MockTemplateService_GetTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTemplateService_GetTemplate_Call) RunAndReturn(run func(string) (*model.Template, error)) *MockTemplateService_GetTemplate_Call {
	_c.Call.Return(run)
	return _c
}

// GetTemplates provides a mock function with given fields:
func (_m *MockTemplateService) GetTemplates() []model.Template {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTemplates")
	}

	var r0 []model.Template
	if rf, ok := ret.Get(0).(func() []model.Template); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Template)
		}
	}

	return r0
}

// MockTemplateService_GetTemplates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplates'
type MockTemplateService_GetTemplates_Call struct {
	*mock.Call
}

// GetTemplates is a helper method to define mock.On call
func (_e *MockTemplateService_Expecter) GetTemplates() *MockTemplateService_GetTemplates_Call {
	return &MockTemplateService_GetTemplates_Call{Call: _e.mock.On("GetTemplates")}
}

func (_c *MockTemplateService_GetTemplates_Call) Run(run func()) *MockTemplateService_GetTemplates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTemplateService_GetTemplates_Call) Return(_a0 []model.Template) *MockTemplateService_GetTemplates_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTemplateService_GetTemplates_Call) RunAndReturn(run func() []model.Template) *MockTemplateService_GetTemplates_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTemplateService creates a new instance of MockTemplateService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTemplateService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTemplateService {
	mock := &MockTemplateService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
