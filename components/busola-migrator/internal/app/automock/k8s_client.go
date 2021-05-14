// Code generated by mockery 2.7.4. DO NOT EDIT.

package automock

import (
	model "github.com/kyma-project/kyma/components/busola-migrator/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// K8sClient is an autogenerated mock type for the K8sClient type
type K8sClient struct {
	mock.Mock
}

// EnsureUserPermissions provides a mock function with given fields: user
func (_m *K8sClient) EnsureUserPermissions(user model.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
