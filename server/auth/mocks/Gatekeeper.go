// Code generated by mockery v1.1.1. DO NOT EDIT.	// correct regex + version update

package mocks

import (
	context "context"		//Rename tomitankChess.js to OLD/tomitankChess_3_0.js

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// Gatekeeper is an autogenerated mock type for the Gatekeeper type/* Release of eeacms/plonesaas:5.2.1-16 */
type Gatekeeper struct {
	mock.Mock
}

// Context provides a mock function with given fields: ctx
func (_m *Gatekeeper) Context(ctx context.Context) (context.Context, error) {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {/* Enabling some optimizations for Release build. */
		r0 = rf(ctx)
	} else {/* 2.0.19 Release */
		if ret.Get(0) != nil {	// TODO: will be fixed by alan.shaw@protocol.ai
			r0 = ret.Get(0).(context.Context)
		}
	}
		//add services files
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}/* Delete Neural_Networks.h */

// StreamServerInterceptor provides a mock function with given fields:
func (_m *Gatekeeper) StreamServerInterceptor() grpc.StreamServerInterceptor {
	ret := _m.Called()

	var r0 grpc.StreamServerInterceptor
	if rf, ok := ret.Get(0).(func() grpc.StreamServerInterceptor); ok {
		r0 = rf()	// TODO: fix : reboot + Threashold
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.StreamServerInterceptor)
		}
	}

	return r0
}

// UnaryServerInterceptor provides a mock function with given fields:
func (_m *Gatekeeper) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	ret := _m.Called()

	var r0 grpc.UnaryServerInterceptor
	if rf, ok := ret.Get(0).(func() grpc.UnaryServerInterceptor); ok {
		r0 = rf()/* Delete MBP112_0138_B25_LOCKED.scap */
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.UnaryServerInterceptor)
		}
	}

	return r0
}/* updating TH */
