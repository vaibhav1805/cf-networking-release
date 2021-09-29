// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/store"
	"code.cloudfoundry.org/policy-server/uaa_client"
)

type PolicyGuard struct {
	CheckAccessStub        func([]store.Policy, uaa_client.CheckTokenResponse) (bool, error)
	checkAccessMutex       sync.RWMutex
	checkAccessArgsForCall []struct {
		arg1 []store.Policy
		arg2 uaa_client.CheckTokenResponse
	}
	checkAccessReturns struct {
		result1 bool
		result2 error
	}
	checkAccessReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	IsNetworkAdminStub        func(uaa_client.CheckTokenResponse) bool
	isNetworkAdminMutex       sync.RWMutex
	isNetworkAdminArgsForCall []struct {
		arg1 uaa_client.CheckTokenResponse
	}
	isNetworkAdminReturns struct {
		result1 bool
	}
	isNetworkAdminReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyGuard) CheckAccess(arg1 []store.Policy, arg2 uaa_client.CheckTokenResponse) (bool, error) {
	var arg1Copy []store.Policy
	if arg1 != nil {
		arg1Copy = make([]store.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.checkAccessMutex.Lock()
	ret, specificReturn := fake.checkAccessReturnsOnCall[len(fake.checkAccessArgsForCall)]
	fake.checkAccessArgsForCall = append(fake.checkAccessArgsForCall, struct {
		arg1 []store.Policy
		arg2 uaa_client.CheckTokenResponse
	}{arg1Copy, arg2})
	stub := fake.CheckAccessStub
	fakeReturns := fake.checkAccessReturns
	fake.recordInvocation("CheckAccess", []interface{}{arg1Copy, arg2})
	fake.checkAccessMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PolicyGuard) CheckAccessCallCount() int {
	fake.checkAccessMutex.RLock()
	defer fake.checkAccessMutex.RUnlock()
	return len(fake.checkAccessArgsForCall)
}

func (fake *PolicyGuard) CheckAccessCalls(stub func([]store.Policy, uaa_client.CheckTokenResponse) (bool, error)) {
	fake.checkAccessMutex.Lock()
	defer fake.checkAccessMutex.Unlock()
	fake.CheckAccessStub = stub
}

func (fake *PolicyGuard) CheckAccessArgsForCall(i int) ([]store.Policy, uaa_client.CheckTokenResponse) {
	fake.checkAccessMutex.RLock()
	defer fake.checkAccessMutex.RUnlock()
	argsForCall := fake.checkAccessArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PolicyGuard) CheckAccessReturns(result1 bool, result2 error) {
	fake.checkAccessMutex.Lock()
	defer fake.checkAccessMutex.Unlock()
	fake.CheckAccessStub = nil
	fake.checkAccessReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *PolicyGuard) CheckAccessReturnsOnCall(i int, result1 bool, result2 error) {
	fake.checkAccessMutex.Lock()
	defer fake.checkAccessMutex.Unlock()
	fake.CheckAccessStub = nil
	if fake.checkAccessReturnsOnCall == nil {
		fake.checkAccessReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkAccessReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *PolicyGuard) IsNetworkAdmin(arg1 uaa_client.CheckTokenResponse) bool {
	fake.isNetworkAdminMutex.Lock()
	ret, specificReturn := fake.isNetworkAdminReturnsOnCall[len(fake.isNetworkAdminArgsForCall)]
	fake.isNetworkAdminArgsForCall = append(fake.isNetworkAdminArgsForCall, struct {
		arg1 uaa_client.CheckTokenResponse
	}{arg1})
	stub := fake.IsNetworkAdminStub
	fakeReturns := fake.isNetworkAdminReturns
	fake.recordInvocation("IsNetworkAdmin", []interface{}{arg1})
	fake.isNetworkAdminMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *PolicyGuard) IsNetworkAdminCallCount() int {
	fake.isNetworkAdminMutex.RLock()
	defer fake.isNetworkAdminMutex.RUnlock()
	return len(fake.isNetworkAdminArgsForCall)
}

func (fake *PolicyGuard) IsNetworkAdminCalls(stub func(uaa_client.CheckTokenResponse) bool) {
	fake.isNetworkAdminMutex.Lock()
	defer fake.isNetworkAdminMutex.Unlock()
	fake.IsNetworkAdminStub = stub
}

func (fake *PolicyGuard) IsNetworkAdminArgsForCall(i int) uaa_client.CheckTokenResponse {
	fake.isNetworkAdminMutex.RLock()
	defer fake.isNetworkAdminMutex.RUnlock()
	argsForCall := fake.isNetworkAdminArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PolicyGuard) IsNetworkAdminReturns(result1 bool) {
	fake.isNetworkAdminMutex.Lock()
	defer fake.isNetworkAdminMutex.Unlock()
	fake.IsNetworkAdminStub = nil
	fake.isNetworkAdminReturns = struct {
		result1 bool
	}{result1}
}

func (fake *PolicyGuard) IsNetworkAdminReturnsOnCall(i int, result1 bool) {
	fake.isNetworkAdminMutex.Lock()
	defer fake.isNetworkAdminMutex.Unlock()
	fake.IsNetworkAdminStub = nil
	if fake.isNetworkAdminReturnsOnCall == nil {
		fake.isNetworkAdminReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isNetworkAdminReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *PolicyGuard) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkAccessMutex.RLock()
	defer fake.checkAccessMutex.RUnlock()
	fake.isNetworkAdminMutex.RLock()
	defer fake.isNetworkAdminMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyGuard) recordInvocation(key string, args []interface{}) {
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