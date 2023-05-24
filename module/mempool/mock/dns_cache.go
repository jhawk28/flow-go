// Code generated by mockery v2.21.4. DO NOT EDIT.

package mempool

import (
	mempool "github.com/onflow/flow-go/module/mempool"
	mock "github.com/stretchr/testify/mock"

	net "net"
)

// DNSCache is an autogenerated mock type for the DNSCache type
type DNSCache struct {
	mock.Mock
}

// GetDomainIp provides a mock function with given fields: _a0
func (_m *DNSCache) GetDomainIp(_a0 string) (*mempool.IpRecord, bool) {
	ret := _m.Called(_a0)

	var r0 *mempool.IpRecord
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (*mempool.IpRecord, bool)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *mempool.IpRecord); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mempool.IpRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetTxtRecord provides a mock function with given fields: _a0
func (_m *DNSCache) GetTxtRecord(_a0 string) (*mempool.TxtRecord, bool) {
	ret := _m.Called(_a0)

	var r0 *mempool.TxtRecord
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (*mempool.TxtRecord, bool)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *mempool.TxtRecord); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mempool.TxtRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// LockIPDomain provides a mock function with given fields: _a0
func (_m *DNSCache) LockIPDomain(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockTxtRecord provides a mock function with given fields: _a0
func (_m *DNSCache) LockTxtRecord(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutIpDomain provides a mock function with given fields: _a0, _a1, _a2
func (_m *DNSCache) PutIpDomain(_a0 string, _a1 []net.IPAddr, _a2 int64) bool {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []net.IPAddr, int64) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PutTxtRecord provides a mock function with given fields: _a0, _a1, _a2
func (_m *DNSCache) PutTxtRecord(_a0 string, _a1 []string, _a2 int64) bool {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, []string, int64) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveIp provides a mock function with given fields: _a0
func (_m *DNSCache) RemoveIp(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveTxt provides a mock function with given fields: _a0
func (_m *DNSCache) RemoveTxt(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Size provides a mock function with given fields:
func (_m *DNSCache) Size() (uint, uint) {
	ret := _m.Called()

	var r0 uint
	var r1 uint
	if rf, ok := ret.Get(0).(func() (uint, uint)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func() uint); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(uint)
	}

	return r0, r1
}

// UpdateIPDomain provides a mock function with given fields: _a0, _a1, _a2
func (_m *DNSCache) UpdateIPDomain(_a0 string, _a1 []net.IPAddr, _a2 int64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []net.IPAddr, int64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxtRecord provides a mock function with given fields: _a0, _a1, _a2
func (_m *DNSCache) UpdateTxtRecord(_a0 string, _a1 []string, _a2 int64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, int64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDNSCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewDNSCache creates a new instance of DNSCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDNSCache(t mockConstructorTestingTNewDNSCache) *DNSCache {
	mock := &DNSCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
