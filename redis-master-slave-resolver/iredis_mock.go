// Code generated by MockGen. DO NOT EDIT.
// Source: iredis.go

// Package main is a generated GoMock package.
package main

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIRedisClientInterface is a mock of IRedisClientInterface interface.
type MockIRedisClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisClientInterfaceMockRecorder
}

// MockIRedisClientInterfaceMockRecorder is the mock recorder for MockIRedisClientInterface.
type MockIRedisClientInterfaceMockRecorder struct {
	mock *MockIRedisClientInterface
}

// NewMockIRedisClientInterface creates a new mock instance.
func NewMockIRedisClientInterface(ctrl *gomock.Controller) *MockIRedisClientInterface {
	mock := &MockIRedisClientInterface{ctrl: ctrl}
	mock.recorder = &MockIRedisClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedisClientInterface) EXPECT() *MockIRedisClientInterfaceMockRecorder {
	return m.recorder
}

// Decr mocks base method.
func (m *MockIRedisClientInterface) Decr(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decr indicates an expected call of Decr.
func (mr *MockIRedisClientInterfaceMockRecorder) Decr(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockIRedisClientInterface)(nil).Decr), ctx, key)
}

// Del mocks base method.
func (m *MockIRedisClientInterface) Del(ctx context.Context, keys ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Del indicates an expected call of Del.
func (mr *MockIRedisClientInterfaceMockRecorder) Del(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockIRedisClientInterface)(nil).Del), varargs...)
}

// Exists mocks base method.
func (m *MockIRedisClientInterface) Exists(ctx context.Context, keys ...string) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exists", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockIRedisClientInterfaceMockRecorder) Exists(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIRedisClientInterface)(nil).Exists), varargs...)
}

// Expire mocks base method.
func (m *MockIRedisClientInterface) Expire(ctx context.Context, key string, expire time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", ctx, key, expire)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expire indicates an expected call of Expire.
func (mr *MockIRedisClientInterfaceMockRecorder) Expire(ctx, key, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockIRedisClientInterface)(nil).Expire), ctx, key, expire)
}

// Get mocks base method.
func (m *MockIRedisClientInterface) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIRedisClientInterfaceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRedisClientInterface)(nil).Get), ctx, key)
}

// Incr mocks base method.
func (m *MockIRedisClientInterface) Incr(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Incr indicates an expected call of Incr.
func (mr *MockIRedisClientInterfaceMockRecorder) Incr(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockIRedisClientInterface)(nil).Incr), ctx, key)
}

// LPush mocks base method.
func (m *MockIRedisClientInterface) LPush(ctx context.Context, key, values string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPush", ctx, key, values)
	ret0, _ := ret[0].(error)
	return ret0
}

// LPush indicates an expected call of LPush.
func (mr *MockIRedisClientInterfaceMockRecorder) LPush(ctx, key, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPush", reflect.TypeOf((*MockIRedisClientInterface)(nil).LPush), ctx, key, values)
}

// LRange mocks base method.
func (m *MockIRedisClientInterface) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRange", ctx, key, start, stop)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LRange indicates an expected call of LRange.
func (mr *MockIRedisClientInterfaceMockRecorder) LRange(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRange", reflect.TypeOf((*MockIRedisClientInterface)(nil).LRange), ctx, key, start, stop)
}

// LRem mocks base method.
func (m *MockIRedisClientInterface) LRem(ctx context.Context, key string, count int64, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRem", ctx, key, count, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// LRem indicates an expected call of LRem.
func (mr *MockIRedisClientInterfaceMockRecorder) LRem(ctx, key, count, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRem", reflect.TypeOf((*MockIRedisClientInterface)(nil).LRem), ctx, key, count, value)
}

// Set mocks base method.
func (m *MockIRedisClientInterface) Set(ctx context.Context, key string, value interface{}, expire time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockIRedisClientInterfaceMockRecorder) Set(ctx, key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockIRedisClientInterface)(nil).Set), ctx, key, value, expire)
}

// SetNX mocks base method.
func (m *MockIRedisClientInterface) SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", ctx, key, value, expire)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX.
func (mr *MockIRedisClientInterfaceMockRecorder) SetNX(ctx, key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockIRedisClientInterface)(nil).SetNX), ctx, key, value, expire)
}

// ZAdd mocks base method.
func (m *MockIRedisClientInterface) ZAdd(ctx context.Context, key string, score float64, member string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAdd", ctx, key, score, member)
	ret0, _ := ret[0].(error)
	return ret0
}

// ZAdd indicates an expected call of ZAdd.
func (mr *MockIRedisClientInterfaceMockRecorder) ZAdd(ctx, key, score, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAdd", reflect.TypeOf((*MockIRedisClientInterface)(nil).ZAdd), ctx, key, score, member)
}

// ZIncrBy mocks base method.
func (m *MockIRedisClientInterface) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrBy", ctx, key, increment, member)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZIncrBy indicates an expected call of ZIncrBy.
func (mr *MockIRedisClientInterfaceMockRecorder) ZIncrBy(ctx, key, increment, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrBy", reflect.TypeOf((*MockIRedisClientInterface)(nil).ZIncrBy), ctx, key, increment, member)
}

// ZScore mocks base method.
func (m *MockIRedisClientInterface) ZScore(ctx context.Context, key, member string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScore", ctx, key, member)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZScore indicates an expected call of ZScore.
func (mr *MockIRedisClientInterfaceMockRecorder) ZScore(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScore", reflect.TypeOf((*MockIRedisClientInterface)(nil).ZScore), ctx, key, member)
}
