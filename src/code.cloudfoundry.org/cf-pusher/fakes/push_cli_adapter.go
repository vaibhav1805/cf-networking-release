// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type PushCLIAdapter struct {
	AppGuidStub        func(string) (string, error)
	appGuidMutex       sync.RWMutex
	appGuidArgsForCall []struct {
		arg1 string
	}
	appGuidReturns struct {
		result1 string
		result2 error
	}
	appGuidReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CheckAppStub        func(string) ([]byte, error)
	checkAppMutex       sync.RWMutex
	checkAppArgsForCall []struct {
		arg1 string
	}
	checkAppReturns struct {
		result1 []byte
		result2 error
	}
	checkAppReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	PushStub        func(string, string, string) error
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	pushReturns struct {
		result1 error
	}
	pushReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PushCLIAdapter) AppGuid(arg1 string) (string, error) {
	fake.appGuidMutex.Lock()
	ret, specificReturn := fake.appGuidReturnsOnCall[len(fake.appGuidArgsForCall)]
	fake.appGuidArgsForCall = append(fake.appGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.AppGuidStub
	fakeReturns := fake.appGuidReturns
	fake.recordInvocation("AppGuid", []interface{}{arg1})
	fake.appGuidMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PushCLIAdapter) AppGuidCallCount() int {
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	return len(fake.appGuidArgsForCall)
}

func (fake *PushCLIAdapter) AppGuidCalls(stub func(string) (string, error)) {
	fake.appGuidMutex.Lock()
	defer fake.appGuidMutex.Unlock()
	fake.AppGuidStub = stub
}

func (fake *PushCLIAdapter) AppGuidArgsForCall(i int) string {
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	argsForCall := fake.appGuidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PushCLIAdapter) AppGuidReturns(result1 string, result2 error) {
	fake.appGuidMutex.Lock()
	defer fake.appGuidMutex.Unlock()
	fake.AppGuidStub = nil
	fake.appGuidReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) AppGuidReturnsOnCall(i int, result1 string, result2 error) {
	fake.appGuidMutex.Lock()
	defer fake.appGuidMutex.Unlock()
	fake.AppGuidStub = nil
	if fake.appGuidReturnsOnCall == nil {
		fake.appGuidReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.appGuidReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) CheckApp(arg1 string) ([]byte, error) {
	fake.checkAppMutex.Lock()
	ret, specificReturn := fake.checkAppReturnsOnCall[len(fake.checkAppArgsForCall)]
	fake.checkAppArgsForCall = append(fake.checkAppArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CheckAppStub
	fakeReturns := fake.checkAppReturns
	fake.recordInvocation("CheckApp", []interface{}{arg1})
	fake.checkAppMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PushCLIAdapter) CheckAppCallCount() int {
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	return len(fake.checkAppArgsForCall)
}

func (fake *PushCLIAdapter) CheckAppCalls(stub func(string) ([]byte, error)) {
	fake.checkAppMutex.Lock()
	defer fake.checkAppMutex.Unlock()
	fake.CheckAppStub = stub
}

func (fake *PushCLIAdapter) CheckAppArgsForCall(i int) string {
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	argsForCall := fake.checkAppArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PushCLIAdapter) CheckAppReturns(result1 []byte, result2 error) {
	fake.checkAppMutex.Lock()
	defer fake.checkAppMutex.Unlock()
	fake.CheckAppStub = nil
	fake.checkAppReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) CheckAppReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.checkAppMutex.Lock()
	defer fake.checkAppMutex.Unlock()
	fake.CheckAppStub = nil
	if fake.checkAppReturnsOnCall == nil {
		fake.checkAppReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.checkAppReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) Push(arg1 string, arg2 string, arg3 string) error {
	fake.pushMutex.Lock()
	ret, specificReturn := fake.pushReturnsOnCall[len(fake.pushArgsForCall)]
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.PushStub
	fakeReturns := fake.pushReturns
	fake.recordInvocation("Push", []interface{}{arg1, arg2, arg3})
	fake.pushMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *PushCLIAdapter) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *PushCLIAdapter) PushCalls(stub func(string, string, string) error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = stub
}

func (fake *PushCLIAdapter) PushArgsForCall(i int) (string, string, string) {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	argsForCall := fake.pushArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PushCLIAdapter) PushReturns(result1 error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 error
	}{result1}
}

func (fake *PushCLIAdapter) PushReturnsOnCall(i int, result1 error) {
	fake.pushMutex.Lock()
	defer fake.pushMutex.Unlock()
	fake.PushStub = nil
	if fake.pushReturnsOnCall == nil {
		fake.pushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PushCLIAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PushCLIAdapter) recordInvocation(key string, args []interface{}) {
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