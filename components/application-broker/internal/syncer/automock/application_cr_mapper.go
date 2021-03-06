// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import internal "github.com/kyma-project/kyma/components/application-broker/internal"
import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"

// ApplicationCRMapper is an autogenerated mock type for the applicationCRMapper type
type ApplicationCRMapper struct {
	mock.Mock
}

// ToModel provides a mock function with given fields: dto
func (_m *ApplicationCRMapper) ToModel(dto *v1alpha1.Application) (*internal.Application, error) {
	ret := _m.Called(dto)

	var r0 *internal.Application
	if rf, ok := ret.Get(0).(func(*v1alpha1.Application) *internal.Application); ok {
		r0 = rf(dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.Application) error); ok {
		r1 = rf(dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
