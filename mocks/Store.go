// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/ppai-plivo/ut/pkg/db"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *Store) GetUserByID(ctx context.Context, id int) (*db.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *db.User
	if rf, ok := ret.Get(0).(func(context.Context, int) *db.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}