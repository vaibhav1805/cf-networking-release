// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/cc_client"
)

type CCClient struct {
	GetAppSpacesStub        func(string, []string) (map[string]string, error)
	getAppSpacesMutex       sync.RWMutex
	getAppSpacesArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	getAppSpacesReturns struct {
		result1 map[string]string
		result2 error
	}
	getAppSpacesReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetLiveAppGUIDsStub        func(string, []string) (map[string]struct{}, error)
	getLiveAppGUIDsMutex       sync.RWMutex
	getLiveAppGUIDsArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	getLiveAppGUIDsReturns struct {
		result1 map[string]struct{}
		result2 error
	}
	getLiveAppGUIDsReturnsOnCall map[int]struct {
		result1 map[string]struct{}
		result2 error
	}
	GetLiveSpaceGUIDsStub        func(string, []string) (map[string]struct{}, error)
	getLiveSpaceGUIDsMutex       sync.RWMutex
	getLiveSpaceGUIDsArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	getLiveSpaceGUIDsReturns struct {
		result1 map[string]struct{}
		result2 error
	}
	getLiveSpaceGUIDsReturnsOnCall map[int]struct {
		result1 map[string]struct{}
		result2 error
	}
	GetSecurityGroupsStub        func(string) ([]cc_client.SecurityGroupResource, error)
	getSecurityGroupsMutex       sync.RWMutex
	getSecurityGroupsArgsForCall []struct {
		arg1 string
	}
	getSecurityGroupsReturns struct {
		result1 []cc_client.SecurityGroupResource
		result2 error
	}
	getSecurityGroupsReturnsOnCall map[int]struct {
		result1 []cc_client.SecurityGroupResource
		result2 error
	}
	GetSpaceStub        func(string, string) (*cc_client.SpaceResponse, error)
	getSpaceMutex       sync.RWMutex
	getSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceReturns struct {
		result1 *cc_client.SpaceResponse
		result2 error
	}
	getSpaceReturnsOnCall map[int]struct {
		result1 *cc_client.SpaceResponse
		result2 error
	}
	GetSpaceGUIDsStub        func(string, []string) ([]string, error)
	getSpaceGUIDsMutex       sync.RWMutex
	getSpaceGUIDsArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	getSpaceGUIDsReturns struct {
		result1 []string
		result2 error
	}
	getSpaceGUIDsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetSubjectSpaceStub        func(string, string, cc_client.SpaceResponse) (*cc_client.SpaceResource, error)
	getSubjectSpaceMutex       sync.RWMutex
	getSubjectSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 cc_client.SpaceResponse
	}
	getSubjectSpaceReturns struct {
		result1 *cc_client.SpaceResource
		result2 error
	}
	getSubjectSpaceReturnsOnCall map[int]struct {
		result1 *cc_client.SpaceResource
		result2 error
	}
	GetSubjectSpacesStub        func(string, string) (map[string]struct{}, error)
	getSubjectSpacesMutex       sync.RWMutex
	getSubjectSpacesArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSubjectSpacesReturns struct {
		result1 map[string]struct{}
		result2 error
	}
	getSubjectSpacesReturnsOnCall map[int]struct {
		result1 map[string]struct{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CCClient) GetAppSpaces(arg1 string, arg2 []string) (map[string]string, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getAppSpacesMutex.Lock()
	ret, specificReturn := fake.getAppSpacesReturnsOnCall[len(fake.getAppSpacesArgsForCall)]
	fake.getAppSpacesArgsForCall = append(fake.getAppSpacesArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.GetAppSpacesStub
	fakeReturns := fake.getAppSpacesReturns
	fake.recordInvocation("GetAppSpaces", []interface{}{arg1, arg2Copy})
	fake.getAppSpacesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetAppSpacesCallCount() int {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	return len(fake.getAppSpacesArgsForCall)
}

func (fake *CCClient) GetAppSpacesCalls(stub func(string, []string) (map[string]string, error)) {
	fake.getAppSpacesMutex.Lock()
	defer fake.getAppSpacesMutex.Unlock()
	fake.GetAppSpacesStub = stub
}

func (fake *CCClient) GetAppSpacesArgsForCall(i int) (string, []string) {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	argsForCall := fake.getAppSpacesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CCClient) GetAppSpacesReturns(result1 map[string]string, result2 error) {
	fake.getAppSpacesMutex.Lock()
	defer fake.getAppSpacesMutex.Unlock()
	fake.GetAppSpacesStub = nil
	fake.getAppSpacesReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetAppSpacesReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.getAppSpacesMutex.Lock()
	defer fake.getAppSpacesMutex.Unlock()
	fake.GetAppSpacesStub = nil
	if fake.getAppSpacesReturnsOnCall == nil {
		fake.getAppSpacesReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.getAppSpacesReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetLiveAppGUIDs(arg1 string, arg2 []string) (map[string]struct{}, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getLiveAppGUIDsMutex.Lock()
	ret, specificReturn := fake.getLiveAppGUIDsReturnsOnCall[len(fake.getLiveAppGUIDsArgsForCall)]
	fake.getLiveAppGUIDsArgsForCall = append(fake.getLiveAppGUIDsArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.GetLiveAppGUIDsStub
	fakeReturns := fake.getLiveAppGUIDsReturns
	fake.recordInvocation("GetLiveAppGUIDs", []interface{}{arg1, arg2Copy})
	fake.getLiveAppGUIDsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetLiveAppGUIDsCallCount() int {
	fake.getLiveAppGUIDsMutex.RLock()
	defer fake.getLiveAppGUIDsMutex.RUnlock()
	return len(fake.getLiveAppGUIDsArgsForCall)
}

func (fake *CCClient) GetLiveAppGUIDsCalls(stub func(string, []string) (map[string]struct{}, error)) {
	fake.getLiveAppGUIDsMutex.Lock()
	defer fake.getLiveAppGUIDsMutex.Unlock()
	fake.GetLiveAppGUIDsStub = stub
}

func (fake *CCClient) GetLiveAppGUIDsArgsForCall(i int) (string, []string) {
	fake.getLiveAppGUIDsMutex.RLock()
	defer fake.getLiveAppGUIDsMutex.RUnlock()
	argsForCall := fake.getLiveAppGUIDsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CCClient) GetLiveAppGUIDsReturns(result1 map[string]struct{}, result2 error) {
	fake.getLiveAppGUIDsMutex.Lock()
	defer fake.getLiveAppGUIDsMutex.Unlock()
	fake.GetLiveAppGUIDsStub = nil
	fake.getLiveAppGUIDsReturns = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetLiveAppGUIDsReturnsOnCall(i int, result1 map[string]struct{}, result2 error) {
	fake.getLiveAppGUIDsMutex.Lock()
	defer fake.getLiveAppGUIDsMutex.Unlock()
	fake.GetLiveAppGUIDsStub = nil
	if fake.getLiveAppGUIDsReturnsOnCall == nil {
		fake.getLiveAppGUIDsReturnsOnCall = make(map[int]struct {
			result1 map[string]struct{}
			result2 error
		})
	}
	fake.getLiveAppGUIDsReturnsOnCall[i] = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetLiveSpaceGUIDs(arg1 string, arg2 []string) (map[string]struct{}, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getLiveSpaceGUIDsMutex.Lock()
	ret, specificReturn := fake.getLiveSpaceGUIDsReturnsOnCall[len(fake.getLiveSpaceGUIDsArgsForCall)]
	fake.getLiveSpaceGUIDsArgsForCall = append(fake.getLiveSpaceGUIDsArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.GetLiveSpaceGUIDsStub
	fakeReturns := fake.getLiveSpaceGUIDsReturns
	fake.recordInvocation("GetLiveSpaceGUIDs", []interface{}{arg1, arg2Copy})
	fake.getLiveSpaceGUIDsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetLiveSpaceGUIDsCallCount() int {
	fake.getLiveSpaceGUIDsMutex.RLock()
	defer fake.getLiveSpaceGUIDsMutex.RUnlock()
	return len(fake.getLiveSpaceGUIDsArgsForCall)
}

func (fake *CCClient) GetLiveSpaceGUIDsCalls(stub func(string, []string) (map[string]struct{}, error)) {
	fake.getLiveSpaceGUIDsMutex.Lock()
	defer fake.getLiveSpaceGUIDsMutex.Unlock()
	fake.GetLiveSpaceGUIDsStub = stub
}

func (fake *CCClient) GetLiveSpaceGUIDsArgsForCall(i int) (string, []string) {
	fake.getLiveSpaceGUIDsMutex.RLock()
	defer fake.getLiveSpaceGUIDsMutex.RUnlock()
	argsForCall := fake.getLiveSpaceGUIDsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CCClient) GetLiveSpaceGUIDsReturns(result1 map[string]struct{}, result2 error) {
	fake.getLiveSpaceGUIDsMutex.Lock()
	defer fake.getLiveSpaceGUIDsMutex.Unlock()
	fake.GetLiveSpaceGUIDsStub = nil
	fake.getLiveSpaceGUIDsReturns = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetLiveSpaceGUIDsReturnsOnCall(i int, result1 map[string]struct{}, result2 error) {
	fake.getLiveSpaceGUIDsMutex.Lock()
	defer fake.getLiveSpaceGUIDsMutex.Unlock()
	fake.GetLiveSpaceGUIDsStub = nil
	if fake.getLiveSpaceGUIDsReturnsOnCall == nil {
		fake.getLiveSpaceGUIDsReturnsOnCall = make(map[int]struct {
			result1 map[string]struct{}
			result2 error
		})
	}
	fake.getLiveSpaceGUIDsReturnsOnCall[i] = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSecurityGroups(arg1 string) ([]cc_client.SecurityGroupResource, error) {
	fake.getSecurityGroupsMutex.Lock()
	ret, specificReturn := fake.getSecurityGroupsReturnsOnCall[len(fake.getSecurityGroupsArgsForCall)]
	fake.getSecurityGroupsArgsForCall = append(fake.getSecurityGroupsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetSecurityGroupsStub
	fakeReturns := fake.getSecurityGroupsReturns
	fake.recordInvocation("GetSecurityGroups", []interface{}{arg1})
	fake.getSecurityGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetSecurityGroupsCallCount() int {
	fake.getSecurityGroupsMutex.RLock()
	defer fake.getSecurityGroupsMutex.RUnlock()
	return len(fake.getSecurityGroupsArgsForCall)
}

func (fake *CCClient) GetSecurityGroupsCalls(stub func(string) ([]cc_client.SecurityGroupResource, error)) {
	fake.getSecurityGroupsMutex.Lock()
	defer fake.getSecurityGroupsMutex.Unlock()
	fake.GetSecurityGroupsStub = stub
}

func (fake *CCClient) GetSecurityGroupsArgsForCall(i int) string {
	fake.getSecurityGroupsMutex.RLock()
	defer fake.getSecurityGroupsMutex.RUnlock()
	argsForCall := fake.getSecurityGroupsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CCClient) GetSecurityGroupsReturns(result1 []cc_client.SecurityGroupResource, result2 error) {
	fake.getSecurityGroupsMutex.Lock()
	defer fake.getSecurityGroupsMutex.Unlock()
	fake.GetSecurityGroupsStub = nil
	fake.getSecurityGroupsReturns = struct {
		result1 []cc_client.SecurityGroupResource
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSecurityGroupsReturnsOnCall(i int, result1 []cc_client.SecurityGroupResource, result2 error) {
	fake.getSecurityGroupsMutex.Lock()
	defer fake.getSecurityGroupsMutex.Unlock()
	fake.GetSecurityGroupsStub = nil
	if fake.getSecurityGroupsReturnsOnCall == nil {
		fake.getSecurityGroupsReturnsOnCall = make(map[int]struct {
			result1 []cc_client.SecurityGroupResource
			result2 error
		})
	}
	fake.getSecurityGroupsReturnsOnCall[i] = struct {
		result1 []cc_client.SecurityGroupResource
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpace(arg1 string, arg2 string) (*cc_client.SpaceResponse, error) {
	fake.getSpaceMutex.Lock()
	ret, specificReturn := fake.getSpaceReturnsOnCall[len(fake.getSpaceArgsForCall)]
	fake.getSpaceArgsForCall = append(fake.getSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetSpaceStub
	fakeReturns := fake.getSpaceReturns
	fake.recordInvocation("GetSpace", []interface{}{arg1, arg2})
	fake.getSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetSpaceCallCount() int {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return len(fake.getSpaceArgsForCall)
}

func (fake *CCClient) GetSpaceCalls(stub func(string, string) (*cc_client.SpaceResponse, error)) {
	fake.getSpaceMutex.Lock()
	defer fake.getSpaceMutex.Unlock()
	fake.GetSpaceStub = stub
}

func (fake *CCClient) GetSpaceArgsForCall(i int) (string, string) {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	argsForCall := fake.getSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CCClient) GetSpaceReturns(result1 *cc_client.SpaceResponse, result2 error) {
	fake.getSpaceMutex.Lock()
	defer fake.getSpaceMutex.Unlock()
	fake.GetSpaceStub = nil
	fake.getSpaceReturns = struct {
		result1 *cc_client.SpaceResponse
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceReturnsOnCall(i int, result1 *cc_client.SpaceResponse, result2 error) {
	fake.getSpaceMutex.Lock()
	defer fake.getSpaceMutex.Unlock()
	fake.GetSpaceStub = nil
	if fake.getSpaceReturnsOnCall == nil {
		fake.getSpaceReturnsOnCall = make(map[int]struct {
			result1 *cc_client.SpaceResponse
			result2 error
		})
	}
	fake.getSpaceReturnsOnCall[i] = struct {
		result1 *cc_client.SpaceResponse
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceGUIDs(arg1 string, arg2 []string) ([]string, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.getSpaceGUIDsMutex.Lock()
	ret, specificReturn := fake.getSpaceGUIDsReturnsOnCall[len(fake.getSpaceGUIDsArgsForCall)]
	fake.getSpaceGUIDsArgsForCall = append(fake.getSpaceGUIDsArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.GetSpaceGUIDsStub
	fakeReturns := fake.getSpaceGUIDsReturns
	fake.recordInvocation("GetSpaceGUIDs", []interface{}{arg1, arg2Copy})
	fake.getSpaceGUIDsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetSpaceGUIDsCallCount() int {
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	return len(fake.getSpaceGUIDsArgsForCall)
}

func (fake *CCClient) GetSpaceGUIDsCalls(stub func(string, []string) ([]string, error)) {
	fake.getSpaceGUIDsMutex.Lock()
	defer fake.getSpaceGUIDsMutex.Unlock()
	fake.GetSpaceGUIDsStub = stub
}

func (fake *CCClient) GetSpaceGUIDsArgsForCall(i int) (string, []string) {
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	argsForCall := fake.getSpaceGUIDsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CCClient) GetSpaceGUIDsReturns(result1 []string, result2 error) {
	fake.getSpaceGUIDsMutex.Lock()
	defer fake.getSpaceGUIDsMutex.Unlock()
	fake.GetSpaceGUIDsStub = nil
	fake.getSpaceGUIDsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceGUIDsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getSpaceGUIDsMutex.Lock()
	defer fake.getSpaceGUIDsMutex.Unlock()
	fake.GetSpaceGUIDsStub = nil
	if fake.getSpaceGUIDsReturnsOnCall == nil {
		fake.getSpaceGUIDsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getSpaceGUIDsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSubjectSpace(arg1 string, arg2 string, arg3 cc_client.SpaceResponse) (*cc_client.SpaceResource, error) {
	fake.getSubjectSpaceMutex.Lock()
	ret, specificReturn := fake.getSubjectSpaceReturnsOnCall[len(fake.getSubjectSpaceArgsForCall)]
	fake.getSubjectSpaceArgsForCall = append(fake.getSubjectSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 cc_client.SpaceResponse
	}{arg1, arg2, arg3})
	stub := fake.GetSubjectSpaceStub
	fakeReturns := fake.getSubjectSpaceReturns
	fake.recordInvocation("GetSubjectSpace", []interface{}{arg1, arg2, arg3})
	fake.getSubjectSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetSubjectSpaceCallCount() int {
	fake.getSubjectSpaceMutex.RLock()
	defer fake.getSubjectSpaceMutex.RUnlock()
	return len(fake.getSubjectSpaceArgsForCall)
}

func (fake *CCClient) GetSubjectSpaceCalls(stub func(string, string, cc_client.SpaceResponse) (*cc_client.SpaceResource, error)) {
	fake.getSubjectSpaceMutex.Lock()
	defer fake.getSubjectSpaceMutex.Unlock()
	fake.GetSubjectSpaceStub = stub
}

func (fake *CCClient) GetSubjectSpaceArgsForCall(i int) (string, string, cc_client.SpaceResponse) {
	fake.getSubjectSpaceMutex.RLock()
	defer fake.getSubjectSpaceMutex.RUnlock()
	argsForCall := fake.getSubjectSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CCClient) GetSubjectSpaceReturns(result1 *cc_client.SpaceResource, result2 error) {
	fake.getSubjectSpaceMutex.Lock()
	defer fake.getSubjectSpaceMutex.Unlock()
	fake.GetSubjectSpaceStub = nil
	fake.getSubjectSpaceReturns = struct {
		result1 *cc_client.SpaceResource
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSubjectSpaceReturnsOnCall(i int, result1 *cc_client.SpaceResource, result2 error) {
	fake.getSubjectSpaceMutex.Lock()
	defer fake.getSubjectSpaceMutex.Unlock()
	fake.GetSubjectSpaceStub = nil
	if fake.getSubjectSpaceReturnsOnCall == nil {
		fake.getSubjectSpaceReturnsOnCall = make(map[int]struct {
			result1 *cc_client.SpaceResource
			result2 error
		})
	}
	fake.getSubjectSpaceReturnsOnCall[i] = struct {
		result1 *cc_client.SpaceResource
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSubjectSpaces(arg1 string, arg2 string) (map[string]struct{}, error) {
	fake.getSubjectSpacesMutex.Lock()
	ret, specificReturn := fake.getSubjectSpacesReturnsOnCall[len(fake.getSubjectSpacesArgsForCall)]
	fake.getSubjectSpacesArgsForCall = append(fake.getSubjectSpacesArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetSubjectSpacesStub
	fakeReturns := fake.getSubjectSpacesReturns
	fake.recordInvocation("GetSubjectSpaces", []interface{}{arg1, arg2})
	fake.getSubjectSpacesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CCClient) GetSubjectSpacesCallCount() int {
	fake.getSubjectSpacesMutex.RLock()
	defer fake.getSubjectSpacesMutex.RUnlock()
	return len(fake.getSubjectSpacesArgsForCall)
}

func (fake *CCClient) GetSubjectSpacesCalls(stub func(string, string) (map[string]struct{}, error)) {
	fake.getSubjectSpacesMutex.Lock()
	defer fake.getSubjectSpacesMutex.Unlock()
	fake.GetSubjectSpacesStub = stub
}

func (fake *CCClient) GetSubjectSpacesArgsForCall(i int) (string, string) {
	fake.getSubjectSpacesMutex.RLock()
	defer fake.getSubjectSpacesMutex.RUnlock()
	argsForCall := fake.getSubjectSpacesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CCClient) GetSubjectSpacesReturns(result1 map[string]struct{}, result2 error) {
	fake.getSubjectSpacesMutex.Lock()
	defer fake.getSubjectSpacesMutex.Unlock()
	fake.GetSubjectSpacesStub = nil
	fake.getSubjectSpacesReturns = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSubjectSpacesReturnsOnCall(i int, result1 map[string]struct{}, result2 error) {
	fake.getSubjectSpacesMutex.Lock()
	defer fake.getSubjectSpacesMutex.Unlock()
	fake.GetSubjectSpacesStub = nil
	if fake.getSubjectSpacesReturnsOnCall == nil {
		fake.getSubjectSpacesReturnsOnCall = make(map[int]struct {
			result1 map[string]struct{}
			result2 error
		})
	}
	fake.getSubjectSpacesReturnsOnCall[i] = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	fake.getLiveAppGUIDsMutex.RLock()
	defer fake.getLiveAppGUIDsMutex.RUnlock()
	fake.getLiveSpaceGUIDsMutex.RLock()
	defer fake.getLiveSpaceGUIDsMutex.RUnlock()
	fake.getSecurityGroupsMutex.RLock()
	defer fake.getSecurityGroupsMutex.RUnlock()
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	fake.getSubjectSpaceMutex.RLock()
	defer fake.getSubjectSpaceMutex.RUnlock()
	fake.getSubjectSpacesMutex.RLock()
	defer fake.getSubjectSpacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CCClient) recordInvocation(key string, args []interface{}) {
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

var _ cc_client.CCClient = new(CCClient)
