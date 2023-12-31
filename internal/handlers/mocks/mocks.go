// Code generated by MockGen. DO NOT EDIT.
// Source: deliveryservices.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeliveryServiceManager is a mock of DeliveryServiceManager interface.
type MockDeliveryServiceManager struct {
	ctrl     *gomock.Controller
	recorder *MockDeliveryServiceManagerMockRecorder
}

// MockDeliveryServiceManagerMockRecorder is the mock recorder for MockDeliveryServiceManager.
type MockDeliveryServiceManagerMockRecorder struct {
	mock *MockDeliveryServiceManager
}

// NewMockDeliveryServiceManager creates a new mock instance.
func NewMockDeliveryServiceManager(ctrl *gomock.Controller) *MockDeliveryServiceManager {
	mock := &MockDeliveryServiceManager{ctrl: ctrl}
	mock.recorder = &MockDeliveryServiceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeliveryServiceManager) EXPECT() *MockDeliveryServiceManagerMockRecorder {
	return m.recorder
}

// DeliveryServicesNearLocation mocks base method.
func (m *MockDeliveryServiceManager) DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliveryServicesNearLocation", ctx, latitude, longitude)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeliveryServicesNearLocation indicates an expected call of DeliveryServicesNearLocation.
func (mr *MockDeliveryServiceManagerMockRecorder) DeliveryServicesNearLocation(ctx, latitude, longitude interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliveryServicesNearLocation", reflect.TypeOf((*MockDeliveryServiceManager)(nil).DeliveryServicesNearLocation), ctx, latitude, longitude)
}
