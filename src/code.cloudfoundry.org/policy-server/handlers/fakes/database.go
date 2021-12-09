// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/cf-networking-helpers/db"
)

type Db struct {
	BeginxStub        func() (db.Transaction, error)
	beginxMutex       sync.RWMutex
	beginxArgsForCall []struct {
	}
	beginxReturns struct {
		result1 db.Transaction
		result2 error
	}
	beginxReturnsOnCall map[int]struct {
		result1 db.Transaction
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Db) Beginx() (db.Transaction, error) {
	fake.beginxMutex.Lock()
	ret, specificReturn := fake.beginxReturnsOnCall[len(fake.beginxArgsForCall)]
	fake.beginxArgsForCall = append(fake.beginxArgsForCall, struct {
	}{})
	stub := fake.BeginxStub
	fakeReturns := fake.beginxReturns
	fake.recordInvocation("Beginx", []interface{}{})
	fake.beginxMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Db) BeginxCallCount() int {
	fake.beginxMutex.RLock()
	defer fake.beginxMutex.RUnlock()
	return len(fake.beginxArgsForCall)
}

func (fake *Db) BeginxCalls(stub func() (db.Transaction, error)) {
	fake.beginxMutex.Lock()
	defer fake.beginxMutex.Unlock()
	fake.BeginxStub = stub
}

func (fake *Db) BeginxReturns(result1 db.Transaction, result2 error) {
	fake.beginxMutex.Lock()
	defer fake.beginxMutex.Unlock()
	fake.BeginxStub = nil
	fake.beginxReturns = struct {
		result1 db.Transaction
		result2 error
	}{result1, result2}
}

func (fake *Db) BeginxReturnsOnCall(i int, result1 db.Transaction, result2 error) {
	fake.beginxMutex.Lock()
	defer fake.beginxMutex.Unlock()
	fake.BeginxStub = nil
	if fake.beginxReturnsOnCall == nil {
		fake.beginxReturnsOnCall = make(map[int]struct {
			result1 db.Transaction
			result2 error
		})
	}
	fake.beginxReturnsOnCall[i] = struct {
		result1 db.Transaction
		result2 error
	}{result1, result2}
}

func (fake *Db) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginxMutex.RLock()
	defer fake.beginxMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Db) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
