// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BucketKeyValueStore is an autogenerated mock type for the BucketKeyValueStore type
type BucketKeyValueStore struct {
	mock.Mock
}

// DeleteAll provides a mock function with given fields: bucket
func (_m *BucketKeyValueStore) DeleteAll(bucket string) error {
	ret := _m.Called(bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: bucket, key
func (_m *BucketKeyValueStore) Get(bucket string, key string) (string, error) {
	ret := _m.Called(bucket, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(bucket, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bucket, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: bucket
func (_m *BucketKeyValueStore) GetAll(bucket string) (map[string]string, error) {
	ret := _m.Called(bucket)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitBucket provides a mock function with given fields: bucket
func (_m *BucketKeyValueStore) InitBucket(bucket string) error {
	ret := _m.Called(bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: bucket, key, value
func (_m *BucketKeyValueStore) Put(bucket string, key string, value string) error {
	ret := _m.Called(bucket, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(bucket, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
