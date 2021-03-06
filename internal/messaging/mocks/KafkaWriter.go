// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import kafka "github.com/segmentio/kafka-go"

import mock "github.com/stretchr/testify/mock"

// KafkaWriter is an autogenerated mock type for the KafkaWriter type
type KafkaWriter struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *KafkaWriter) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteMessages provides a mock function with given fields: _a0, _a1
func (_m *KafkaWriter) WriteMessages(_a0 context.Context, _a1 ...kafka.Message) error {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...kafka.Message) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
