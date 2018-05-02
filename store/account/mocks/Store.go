// Code generated by mockery v1.0.0
package mocks

import common "github.com/ethereum/go-ethereum/common"
import mock "github.com/stretchr/testify/mock"
import model "github.com/maichain/eth-indexer/model"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// DeleteAccounts provides a mock function with given fields: fromBlock
func (_m *Store) DeleteAccounts(fromBlock int64) error {
	ret := _m.Called(fromBlock)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(fromBlock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteContractCodes provides a mock function with given fields: fromBlock
func (_m *Store) DeleteContractCodes(fromBlock int64) error {
	ret := _m.Called(fromBlock)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(fromBlock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteContracts provides a mock function with given fields: fromBlock
func (_m *Store) DeleteContracts(fromBlock int64) error {
	ret := _m.Called(fromBlock)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(fromBlock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStateBlocks provides a mock function with given fields: fromBlock
func (_m *Store) DeleteStateBlocks(fromBlock int64) error {
	ret := _m.Called(fromBlock)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(fromBlock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAccount provides a mock function with given fields: address, blockNr
func (_m *Store) FindAccount(address common.Address, blockNr ...int64) (*model.Account, error) {
	_va := make([]interface{}, len(blockNr))
	for _i := range blockNr {
		_va[_i] = blockNr[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(common.Address, ...int64) *model.Account); ok {
		r0 = rf(address, blockNr...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, ...int64) error); ok {
		r1 = rf(address, blockNr...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindContract provides a mock function with given fields: address, blockNr
func (_m *Store) FindContract(address common.Address, blockNr ...int64) (*model.Contract, error) {
	_va := make([]interface{}, len(blockNr))
	for _i := range blockNr {
		_va[_i] = blockNr[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.Contract
	if rf, ok := ret.Get(0).(func(common.Address, ...int64) *model.Contract); ok {
		r0 = rf(address, blockNr...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Contract)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, ...int64) error); ok {
		r1 = rf(address, blockNr...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindContractCode provides a mock function with given fields: address
func (_m *Store) FindContractCode(address common.Address) (*model.ContractCode, error) {
	ret := _m.Called(address)

	var r0 *model.ContractCode
	if rf, ok := ret.Get(0).(func(common.Address) *model.ContractCode); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ContractCode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindStateBlock provides a mock function with given fields: blockNr
func (_m *Store) FindStateBlock(blockNr int64) (*model.StateBlock, error) {
	ret := _m.Called(blockNr)

	var r0 *model.StateBlock
	if rf, ok := ret.Get(0).(func(int64) *model.StateBlock); ok {
		r0 = rf(blockNr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StateBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(blockNr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertAccount provides a mock function with given fields: _a0
func (_m *Store) InsertAccount(_a0 *model.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertContract provides a mock function with given fields: contract
func (_m *Store) InsertContract(contract *model.Contract) error {
	ret := _m.Called(contract)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Contract) error); ok {
		r0 = rf(contract)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertContractCode provides a mock function with given fields: code
func (_m *Store) InsertContractCode(code *model.ContractCode) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ContractCode) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertStateBlock provides a mock function with given fields: block
func (_m *Store) InsertStateBlock(block *model.StateBlock) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.StateBlock) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LastStateBlock provides a mock function with given fields:
func (_m *Store) LastStateBlock() (*model.StateBlock, error) {
	ret := _m.Called()

	var r0 *model.StateBlock
	if rf, ok := ret.Get(0).(func() *model.StateBlock); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StateBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}