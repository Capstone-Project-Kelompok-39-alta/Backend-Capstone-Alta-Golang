// Code generated by mockery v2.12.2. DO NOT EDIT.

package send_customer

import (
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	mock "github.com/stretchr/testify/mock"
)

// SendCustomerRepository is an autogenerated mock type for the SendCustomerRepository type
type SendCustomerRepository struct {
	mock.Mock
}

// SendEmail provides a mock function with given fields: message
func (_m *SendCustomerRepository) SendEmail(message entities.SendCustomer) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.SendCustomer) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
