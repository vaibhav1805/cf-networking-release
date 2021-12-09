// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/api"
	"code.cloudfoundry.org/policy-server/store"
)

type PolicyMapper struct {
	AsBytesStub        func([]store.Policy) ([]byte, error)
	asBytesMutex       sync.RWMutex
	asBytesArgsForCall []struct {
		arg1 []store.Policy
	}
	asBytesReturns struct {
		result1 []byte
		result2 error
	}
	asBytesReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	AsStorePolicyStub        func([]byte) ([]store.Policy, error)
	asStorePolicyMutex       sync.RWMutex
	asStorePolicyArgsForCall []struct {
		arg1 []byte
	}
	asStorePolicyReturns struct {
		result1 []store.Policy
		result2 error
	}
	asStorePolicyReturnsOnCall map[int]struct {
		result1 []store.Policy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyMapper) AsBytes(arg1 []store.Policy) ([]byte, error) {
	var arg1Copy []store.Policy
	if arg1 != nil {
		arg1Copy = make([]store.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.asBytesMutex.Lock()
	ret, specificReturn := fake.asBytesReturnsOnCall[len(fake.asBytesArgsForCall)]
	fake.asBytesArgsForCall = append(fake.asBytesArgsForCall, struct {
		arg1 []store.Policy
	}{arg1Copy})
	stub := fake.AsBytesStub
	fakeReturns := fake.asBytesReturns
	fake.recordInvocation("AsBytes", []interface{}{arg1Copy})
	fake.asBytesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PolicyMapper) AsBytesCallCount() int {
	fake.asBytesMutex.RLock()
	defer fake.asBytesMutex.RUnlock()
	return len(fake.asBytesArgsForCall)
}

func (fake *PolicyMapper) AsBytesCalls(stub func([]store.Policy) ([]byte, error)) {
	fake.asBytesMutex.Lock()
	defer fake.asBytesMutex.Unlock()
	fake.AsBytesStub = stub
}

func (fake *PolicyMapper) AsBytesArgsForCall(i int) []store.Policy {
	fake.asBytesMutex.RLock()
	defer fake.asBytesMutex.RUnlock()
	argsForCall := fake.asBytesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PolicyMapper) AsBytesReturns(result1 []byte, result2 error) {
	fake.asBytesMutex.Lock()
	defer fake.asBytesMutex.Unlock()
	fake.AsBytesStub = nil
	fake.asBytesReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) AsBytesReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.asBytesMutex.Lock()
	defer fake.asBytesMutex.Unlock()
	fake.AsBytesStub = nil
	if fake.asBytesReturnsOnCall == nil {
		fake.asBytesReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.asBytesReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) AsStorePolicy(arg1 []byte) ([]store.Policy, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.asStorePolicyMutex.Lock()
	ret, specificReturn := fake.asStorePolicyReturnsOnCall[len(fake.asStorePolicyArgsForCall)]
	fake.asStorePolicyArgsForCall = append(fake.asStorePolicyArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.AsStorePolicyStub
	fakeReturns := fake.asStorePolicyReturns
	fake.recordInvocation("AsStorePolicy", []interface{}{arg1Copy})
	fake.asStorePolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PolicyMapper) AsStorePolicyCallCount() int {
	fake.asStorePolicyMutex.RLock()
	defer fake.asStorePolicyMutex.RUnlock()
	return len(fake.asStorePolicyArgsForCall)
}

func (fake *PolicyMapper) AsStorePolicyCalls(stub func([]byte) ([]store.Policy, error)) {
	fake.asStorePolicyMutex.Lock()
	defer fake.asStorePolicyMutex.Unlock()
	fake.AsStorePolicyStub = stub
}

func (fake *PolicyMapper) AsStorePolicyArgsForCall(i int) []byte {
	fake.asStorePolicyMutex.RLock()
	defer fake.asStorePolicyMutex.RUnlock()
	argsForCall := fake.asStorePolicyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PolicyMapper) AsStorePolicyReturns(result1 []store.Policy, result2 error) {
	fake.asStorePolicyMutex.Lock()
	defer fake.asStorePolicyMutex.Unlock()
	fake.AsStorePolicyStub = nil
	fake.asStorePolicyReturns = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) AsStorePolicyReturnsOnCall(i int, result1 []store.Policy, result2 error) {
	fake.asStorePolicyMutex.Lock()
	defer fake.asStorePolicyMutex.Unlock()
	fake.AsStorePolicyStub = nil
	if fake.asStorePolicyReturnsOnCall == nil {
		fake.asStorePolicyReturnsOnCall = make(map[int]struct {
			result1 []store.Policy
			result2 error
		})
	}
	fake.asStorePolicyReturnsOnCall[i] = struct {
		result1 []store.Policy
		result2 error
	}{result1, result2}
}

func (fake *PolicyMapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.asBytesMutex.RLock()
	defer fake.asBytesMutex.RUnlock()
	fake.asStorePolicyMutex.RLock()
	defer fake.asStorePolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyMapper) recordInvocation(key string, args []interface{}) {
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

var _ api.PolicyMapper = new(PolicyMapper)
