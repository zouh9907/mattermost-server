// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// ExtRefStore is an autogenerated mock type for the ExtRefStore type
type ExtRefStore struct {
	mock.Mock
}

// GetByAliasUserId provides a mock function with given fields: aliasUserId
func (_m *ExtRefStore) GetByAliasUserId(aliasUserId string) (*model.ExtRef, error) {
	ret := _m.Called(aliasUserId)

	var r0 *model.ExtRef
	if rf, ok := ret.Get(0).(func(string) *model.ExtRef); ok {
		r0 = rf(aliasUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExtRef)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(aliasUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByExtIdAndPlatform provides a mock function with given fields: externalId, externalPlatform
func (_m *ExtRefStore) GetByExtIdAndPlatform(externalId string, externalPlatform string) (*model.ExtRef, error) {
	ret := _m.Called(externalId, externalPlatform)

	var r0 *model.ExtRef
	if rf, ok := ret.Get(0).(func(string, string) *model.ExtRef); ok {
		r0 = rf(externalId, externalPlatform)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExtRef)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(externalId, externalPlatform)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRealUserIdAndPlatform provides a mock function with given fields: realUserId, externalPlatform
func (_m *ExtRefStore) GetByRealUserIdAndPlatform(realUserId string, externalPlatform string) (*model.ExtRef, error) {
	ret := _m.Called(realUserId, externalPlatform)

	var r0 *model.ExtRef
	if rf, ok := ret.Get(0).(func(string, string) *model.ExtRef); ok {
		r0 = rf(realUserId, externalPlatform)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExtRef)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(realUserId, externalPlatform)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ext_ref
func (_m *ExtRefStore) Save(ext_ref *model.ExtRef) (*model.ExtRef, error) {
	ret := _m.Called(ext_ref)

	var r0 *model.ExtRef
	if rf, ok := ret.Get(0).(func(*model.ExtRef) *model.ExtRef); ok {
		r0 = rf(ext_ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExtRef)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ExtRef) error); ok {
		r1 = rf(ext_ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlink provides a mock function with given fields: realUserId, externalPlatform
func (_m *ExtRefStore) Unlink(realUserId string, externalPlatform string) error {
	ret := _m.Called(realUserId, externalPlatform)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(realUserId, externalPlatform)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRealId provides a mock function with given fields: realUserId, externalId, externalPlatform
func (_m *ExtRefStore) UpdateRealId(realUserId string, externalId string, externalPlatform string) error {
	ret := _m.Called(realUserId, externalId, externalPlatform)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(realUserId, externalId, externalPlatform)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}