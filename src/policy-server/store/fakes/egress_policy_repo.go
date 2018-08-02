// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/db"
	"policy-server/store"
	"sync"
)

type EgressPolicyRepo struct {
	CreateTerminalStub        func(tx db.Transaction) (int64, error)
	createTerminalMutex       sync.RWMutex
	createTerminalArgsForCall []struct {
		tx db.Transaction
	}
	createTerminalReturns struct {
		result1 int64
		result2 error
	}
	createTerminalReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateAppStub        func(tx db.Transaction, sourceTerminalID int64, appGUID string) (int64, error)
	createAppMutex       sync.RWMutex
	createAppArgsForCall []struct {
		tx               db.Transaction
		sourceTerminalID int64
		appGUID          string
	}
	createAppReturns struct {
		result1 int64
		result2 error
	}
	createAppReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateIPRangeStub        func(tx db.Transaction, destinationTerminalID int64, startIP, endIP, protocol string, startPort, endPort int64) (int64, error)
	createIPRangeMutex       sync.RWMutex
	createIPRangeArgsForCall []struct {
		tx                    db.Transaction
		destinationTerminalID int64
		startIP               string
		endIP                 string
		protocol              string
		startPort             int64
		endPort               int64
	}
	createIPRangeReturns struct {
		result1 int64
		result2 error
	}
	createIPRangeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateEgressPolicyStub        func(tx db.Transaction, sourceTerminalID, destinationTerminalID int64) (int64, error)
	createEgressPolicyMutex       sync.RWMutex
	createEgressPolicyArgsForCall []struct {
		tx                    db.Transaction
		sourceTerminalID      int64
		destinationTerminalID int64
	}
	createEgressPolicyReturns struct {
		result1 int64
		result2 error
	}
	createEgressPolicyReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetTerminalByAppGUIDStub        func(tx db.Transaction, appGUID string) (int64, error)
	getTerminalByAppGUIDMutex       sync.RWMutex
	getTerminalByAppGUIDArgsForCall []struct {
		tx      db.Transaction
		appGUID string
	}
	getTerminalByAppGUIDReturns struct {
		result1 int64
		result2 error
	}
	getTerminalByAppGUIDReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetAllPoliciesStub        func() ([]store.EgressPolicy, error)
	getAllPoliciesMutex       sync.RWMutex
	getAllPoliciesArgsForCall []struct{}
	getAllPoliciesReturns     struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getAllPoliciesReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetByGuidsStub        func(ids []string) ([]store.EgressPolicy, error)
	getByGuidsMutex       sync.RWMutex
	getByGuidsArgsForCall []struct {
		ids []string
	}
	getByGuidsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getByGuidsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetIDsByEgressPolicyStub        func(tx db.Transaction, egressPolicy store.EgressPolicy) (store.EgressPolicyIDCollection, error)
	getIDsByEgressPolicyMutex       sync.RWMutex
	getIDsByEgressPolicyArgsForCall []struct {
		tx           db.Transaction
		egressPolicy store.EgressPolicy
	}
	getIDsByEgressPolicyReturns struct {
		result1 store.EgressPolicyIDCollection
		result2 error
	}
	getIDsByEgressPolicyReturnsOnCall map[int]struct {
		result1 store.EgressPolicyIDCollection
		result2 error
	}
	DeleteEgressPolicyStub        func(tx db.Transaction, egressPolicyID int64) error
	deleteEgressPolicyMutex       sync.RWMutex
	deleteEgressPolicyArgsForCall []struct {
		tx             db.Transaction
		egressPolicyID int64
	}
	deleteEgressPolicyReturns struct {
		result1 error
	}
	deleteEgressPolicyReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteIPRangeStub        func(tx db.Transaction, ipRangeID int64) error
	deleteIPRangeMutex       sync.RWMutex
	deleteIPRangeArgsForCall []struct {
		tx        db.Transaction
		ipRangeID int64
	}
	deleteIPRangeReturns struct {
		result1 error
	}
	deleteIPRangeReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteTerminalStub        func(tx db.Transaction, terminalID int64) error
	deleteTerminalMutex       sync.RWMutex
	deleteTerminalArgsForCall []struct {
		tx         db.Transaction
		terminalID int64
	}
	deleteTerminalReturns struct {
		result1 error
	}
	deleteTerminalReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteAppStub        func(tx db.Transaction, appID int64) error
	deleteAppMutex       sync.RWMutex
	deleteAppArgsForCall []struct {
		tx    db.Transaction
		appID int64
	}
	deleteAppReturns struct {
		result1 error
	}
	deleteAppReturnsOnCall map[int]struct {
		result1 error
	}
	IsTerminalInUseStub        func(tx db.Transaction, terminalID int64) (bool, error)
	isTerminalInUseMutex       sync.RWMutex
	isTerminalInUseArgsForCall []struct {
		tx         db.Transaction
		terminalID int64
	}
	isTerminalInUseReturns struct {
		result1 bool
		result2 error
	}
	isTerminalInUseReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyRepo) CreateTerminal(tx db.Transaction) (int64, error) {
	fake.createTerminalMutex.Lock()
	ret, specificReturn := fake.createTerminalReturnsOnCall[len(fake.createTerminalArgsForCall)]
	fake.createTerminalArgsForCall = append(fake.createTerminalArgsForCall, struct {
		tx db.Transaction
	}{tx})
	fake.recordInvocation("CreateTerminal", []interface{}{tx})
	fake.createTerminalMutex.Unlock()
	if fake.CreateTerminalStub != nil {
		return fake.CreateTerminalStub(tx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createTerminalReturns.result1, fake.createTerminalReturns.result2
}

func (fake *EgressPolicyRepo) CreateTerminalCallCount() int {
	fake.createTerminalMutex.RLock()
	defer fake.createTerminalMutex.RUnlock()
	return len(fake.createTerminalArgsForCall)
}

func (fake *EgressPolicyRepo) CreateTerminalArgsForCall(i int) db.Transaction {
	fake.createTerminalMutex.RLock()
	defer fake.createTerminalMutex.RUnlock()
	return fake.createTerminalArgsForCall[i].tx
}

func (fake *EgressPolicyRepo) CreateTerminalReturns(result1 int64, result2 error) {
	fake.CreateTerminalStub = nil
	fake.createTerminalReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateTerminalReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateTerminalStub = nil
	if fake.createTerminalReturnsOnCall == nil {
		fake.createTerminalReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createTerminalReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateApp(tx db.Transaction, sourceTerminalID int64, appGUID string) (int64, error) {
	fake.createAppMutex.Lock()
	ret, specificReturn := fake.createAppReturnsOnCall[len(fake.createAppArgsForCall)]
	fake.createAppArgsForCall = append(fake.createAppArgsForCall, struct {
		tx               db.Transaction
		sourceTerminalID int64
		appGUID          string
	}{tx, sourceTerminalID, appGUID})
	fake.recordInvocation("CreateApp", []interface{}{tx, sourceTerminalID, appGUID})
	fake.createAppMutex.Unlock()
	if fake.CreateAppStub != nil {
		return fake.CreateAppStub(tx, sourceTerminalID, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createAppReturns.result1, fake.createAppReturns.result2
}

func (fake *EgressPolicyRepo) CreateAppCallCount() int {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return len(fake.createAppArgsForCall)
}

func (fake *EgressPolicyRepo) CreateAppArgsForCall(i int) (db.Transaction, int64, string) {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return fake.createAppArgsForCall[i].tx, fake.createAppArgsForCall[i].sourceTerminalID, fake.createAppArgsForCall[i].appGUID
}

func (fake *EgressPolicyRepo) CreateAppReturns(result1 int64, result2 error) {
	fake.CreateAppStub = nil
	fake.createAppReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateAppReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateAppStub = nil
	if fake.createAppReturnsOnCall == nil {
		fake.createAppReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createAppReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateIPRange(tx db.Transaction, destinationTerminalID int64, startIP string, endIP string, protocol string, startPort int64, endPort int64) (int64, error) {
	fake.createIPRangeMutex.Lock()
	ret, specificReturn := fake.createIPRangeReturnsOnCall[len(fake.createIPRangeArgsForCall)]
	fake.createIPRangeArgsForCall = append(fake.createIPRangeArgsForCall, struct {
		tx                    db.Transaction
		destinationTerminalID int64
		startIP               string
		endIP                 string
		protocol              string
		startPort             int64
		endPort               int64
	}{tx, destinationTerminalID, startIP, endIP, protocol, startPort, endPort})
	fake.recordInvocation("CreateIPRange", []interface{}{tx, destinationTerminalID, startIP, endIP, protocol, startPort, endPort})
	fake.createIPRangeMutex.Unlock()
	if fake.CreateIPRangeStub != nil {
		return fake.CreateIPRangeStub(tx, destinationTerminalID, startIP, endIP, protocol, startPort, endPort)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createIPRangeReturns.result1, fake.createIPRangeReturns.result2
}

func (fake *EgressPolicyRepo) CreateIPRangeCallCount() int {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return len(fake.createIPRangeArgsForCall)
}

func (fake *EgressPolicyRepo) CreateIPRangeArgsForCall(i int) (db.Transaction, int64, string, string, string, int64, int64) {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return fake.createIPRangeArgsForCall[i].tx, fake.createIPRangeArgsForCall[i].destinationTerminalID, fake.createIPRangeArgsForCall[i].startIP, fake.createIPRangeArgsForCall[i].endIP, fake.createIPRangeArgsForCall[i].protocol, fake.createIPRangeArgsForCall[i].startPort, fake.createIPRangeArgsForCall[i].endPort
}

func (fake *EgressPolicyRepo) CreateIPRangeReturns(result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	fake.createIPRangeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateIPRangeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	if fake.createIPRangeReturnsOnCall == nil {
		fake.createIPRangeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createIPRangeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateEgressPolicy(tx db.Transaction, sourceTerminalID int64, destinationTerminalID int64) (int64, error) {
	fake.createEgressPolicyMutex.Lock()
	ret, specificReturn := fake.createEgressPolicyReturnsOnCall[len(fake.createEgressPolicyArgsForCall)]
	fake.createEgressPolicyArgsForCall = append(fake.createEgressPolicyArgsForCall, struct {
		tx                    db.Transaction
		sourceTerminalID      int64
		destinationTerminalID int64
	}{tx, sourceTerminalID, destinationTerminalID})
	fake.recordInvocation("CreateEgressPolicy", []interface{}{tx, sourceTerminalID, destinationTerminalID})
	fake.createEgressPolicyMutex.Unlock()
	if fake.CreateEgressPolicyStub != nil {
		return fake.CreateEgressPolicyStub(tx, sourceTerminalID, destinationTerminalID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createEgressPolicyReturns.result1, fake.createEgressPolicyReturns.result2
}

func (fake *EgressPolicyRepo) CreateEgressPolicyCallCount() int {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	return len(fake.createEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) CreateEgressPolicyArgsForCall(i int) (db.Transaction, int64, int64) {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	return fake.createEgressPolicyArgsForCall[i].tx, fake.createEgressPolicyArgsForCall[i].sourceTerminalID, fake.createEgressPolicyArgsForCall[i].destinationTerminalID
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturns(result1 int64, result2 error) {
	fake.CreateEgressPolicyStub = nil
	fake.createEgressPolicyReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateEgressPolicyStub = nil
	if fake.createEgressPolicyReturnsOnCall == nil {
		fake.createEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createEgressPolicyReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUID(tx db.Transaction, appGUID string) (int64, error) {
	fake.getTerminalByAppGUIDMutex.Lock()
	ret, specificReturn := fake.getTerminalByAppGUIDReturnsOnCall[len(fake.getTerminalByAppGUIDArgsForCall)]
	fake.getTerminalByAppGUIDArgsForCall = append(fake.getTerminalByAppGUIDArgsForCall, struct {
		tx      db.Transaction
		appGUID string
	}{tx, appGUID})
	fake.recordInvocation("GetTerminalByAppGUID", []interface{}{tx, appGUID})
	fake.getTerminalByAppGUIDMutex.Unlock()
	if fake.GetTerminalByAppGUIDStub != nil {
		return fake.GetTerminalByAppGUIDStub(tx, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTerminalByAppGUIDReturns.result1, fake.getTerminalByAppGUIDReturns.result2
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDCallCount() int {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	return len(fake.getTerminalByAppGUIDArgsForCall)
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDArgsForCall(i int) (db.Transaction, string) {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	return fake.getTerminalByAppGUIDArgsForCall[i].tx, fake.getTerminalByAppGUIDArgsForCall[i].appGUID
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturns(result1 int64, result2 error) {
	fake.GetTerminalByAppGUIDStub = nil
	fake.getTerminalByAppGUIDReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturnsOnCall(i int, result1 int64, result2 error) {
	fake.GetTerminalByAppGUIDStub = nil
	if fake.getTerminalByAppGUIDReturnsOnCall == nil {
		fake.getTerminalByAppGUIDReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.getTerminalByAppGUIDReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetAllPolicies() ([]store.EgressPolicy, error) {
	fake.getAllPoliciesMutex.Lock()
	ret, specificReturn := fake.getAllPoliciesReturnsOnCall[len(fake.getAllPoliciesArgsForCall)]
	fake.getAllPoliciesArgsForCall = append(fake.getAllPoliciesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllPolicies", []interface{}{})
	fake.getAllPoliciesMutex.Unlock()
	if fake.GetAllPoliciesStub != nil {
		return fake.GetAllPoliciesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAllPoliciesReturns.result1, fake.getAllPoliciesReturns.result2
}

func (fake *EgressPolicyRepo) GetAllPoliciesCallCount() int {
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	return len(fake.getAllPoliciesArgsForCall)
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetAllPoliciesStub = nil
	fake.getAllPoliciesReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetAllPoliciesStub = nil
	if fake.getAllPoliciesReturnsOnCall == nil {
		fake.getAllPoliciesReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getAllPoliciesReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGuids(ids []string) ([]store.EgressPolicy, error) {
	var idsCopy []string
	if ids != nil {
		idsCopy = make([]string, len(ids))
		copy(idsCopy, ids)
	}
	fake.getByGuidsMutex.Lock()
	ret, specificReturn := fake.getByGuidsReturnsOnCall[len(fake.getByGuidsArgsForCall)]
	fake.getByGuidsArgsForCall = append(fake.getByGuidsArgsForCall, struct {
		ids []string
	}{idsCopy})
	fake.recordInvocation("GetByGuids", []interface{}{idsCopy})
	fake.getByGuidsMutex.Unlock()
	if fake.GetByGuidsStub != nil {
		return fake.GetByGuidsStub(ids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByGuidsReturns.result1, fake.getByGuidsReturns.result2
}

func (fake *EgressPolicyRepo) GetByGuidsCallCount() int {
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	return len(fake.getByGuidsArgsForCall)
}

func (fake *EgressPolicyRepo) GetByGuidsArgsForCall(i int) []string {
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	return fake.getByGuidsArgsForCall[i].ids
}

func (fake *EgressPolicyRepo) GetByGuidsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetByGuidsStub = nil
	fake.getByGuidsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGuidsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetByGuidsStub = nil
	if fake.getByGuidsReturnsOnCall == nil {
		fake.getByGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getByGuidsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetIDsByEgressPolicy(tx db.Transaction, egressPolicy store.EgressPolicy) (store.EgressPolicyIDCollection, error) {
	fake.getIDsByEgressPolicyMutex.Lock()
	ret, specificReturn := fake.getIDsByEgressPolicyReturnsOnCall[len(fake.getIDsByEgressPolicyArgsForCall)]
	fake.getIDsByEgressPolicyArgsForCall = append(fake.getIDsByEgressPolicyArgsForCall, struct {
		tx           db.Transaction
		egressPolicy store.EgressPolicy
	}{tx, egressPolicy})
	fake.recordInvocation("GetIDsByEgressPolicy", []interface{}{tx, egressPolicy})
	fake.getIDsByEgressPolicyMutex.Unlock()
	if fake.GetIDsByEgressPolicyStub != nil {
		return fake.GetIDsByEgressPolicyStub(tx, egressPolicy)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getIDsByEgressPolicyReturns.result1, fake.getIDsByEgressPolicyReturns.result2
}

func (fake *EgressPolicyRepo) GetIDsByEgressPolicyCallCount() int {
	fake.getIDsByEgressPolicyMutex.RLock()
	defer fake.getIDsByEgressPolicyMutex.RUnlock()
	return len(fake.getIDsByEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) GetIDsByEgressPolicyArgsForCall(i int) (db.Transaction, store.EgressPolicy) {
	fake.getIDsByEgressPolicyMutex.RLock()
	defer fake.getIDsByEgressPolicyMutex.RUnlock()
	return fake.getIDsByEgressPolicyArgsForCall[i].tx, fake.getIDsByEgressPolicyArgsForCall[i].egressPolicy
}

func (fake *EgressPolicyRepo) GetIDsByEgressPolicyReturns(result1 store.EgressPolicyIDCollection, result2 error) {
	fake.GetIDsByEgressPolicyStub = nil
	fake.getIDsByEgressPolicyReturns = struct {
		result1 store.EgressPolicyIDCollection
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetIDsByEgressPolicyReturnsOnCall(i int, result1 store.EgressPolicyIDCollection, result2 error) {
	fake.GetIDsByEgressPolicyStub = nil
	if fake.getIDsByEgressPolicyReturnsOnCall == nil {
		fake.getIDsByEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 store.EgressPolicyIDCollection
			result2 error
		})
	}
	fake.getIDsByEgressPolicyReturnsOnCall[i] = struct {
		result1 store.EgressPolicyIDCollection
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) DeleteEgressPolicy(tx db.Transaction, egressPolicyID int64) error {
	fake.deleteEgressPolicyMutex.Lock()
	ret, specificReturn := fake.deleteEgressPolicyReturnsOnCall[len(fake.deleteEgressPolicyArgsForCall)]
	fake.deleteEgressPolicyArgsForCall = append(fake.deleteEgressPolicyArgsForCall, struct {
		tx             db.Transaction
		egressPolicyID int64
	}{tx, egressPolicyID})
	fake.recordInvocation("DeleteEgressPolicy", []interface{}{tx, egressPolicyID})
	fake.deleteEgressPolicyMutex.Unlock()
	if fake.DeleteEgressPolicyStub != nil {
		return fake.DeleteEgressPolicyStub(tx, egressPolicyID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteEgressPolicyReturns.result1
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyCallCount() int {
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	return len(fake.deleteEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyArgsForCall(i int) (db.Transaction, int64) {
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	return fake.deleteEgressPolicyArgsForCall[i].tx, fake.deleteEgressPolicyArgsForCall[i].egressPolicyID
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyReturns(result1 error) {
	fake.DeleteEgressPolicyStub = nil
	fake.deleteEgressPolicyReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyReturnsOnCall(i int, result1 error) {
	fake.DeleteEgressPolicyStub = nil
	if fake.deleteEgressPolicyReturnsOnCall == nil {
		fake.deleteEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteEgressPolicyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteIPRange(tx db.Transaction, ipRangeID int64) error {
	fake.deleteIPRangeMutex.Lock()
	ret, specificReturn := fake.deleteIPRangeReturnsOnCall[len(fake.deleteIPRangeArgsForCall)]
	fake.deleteIPRangeArgsForCall = append(fake.deleteIPRangeArgsForCall, struct {
		tx        db.Transaction
		ipRangeID int64
	}{tx, ipRangeID})
	fake.recordInvocation("DeleteIPRange", []interface{}{tx, ipRangeID})
	fake.deleteIPRangeMutex.Unlock()
	if fake.DeleteIPRangeStub != nil {
		return fake.DeleteIPRangeStub(tx, ipRangeID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteIPRangeReturns.result1
}

func (fake *EgressPolicyRepo) DeleteIPRangeCallCount() int {
	fake.deleteIPRangeMutex.RLock()
	defer fake.deleteIPRangeMutex.RUnlock()
	return len(fake.deleteIPRangeArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteIPRangeArgsForCall(i int) (db.Transaction, int64) {
	fake.deleteIPRangeMutex.RLock()
	defer fake.deleteIPRangeMutex.RUnlock()
	return fake.deleteIPRangeArgsForCall[i].tx, fake.deleteIPRangeArgsForCall[i].ipRangeID
}

func (fake *EgressPolicyRepo) DeleteIPRangeReturns(result1 error) {
	fake.DeleteIPRangeStub = nil
	fake.deleteIPRangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteIPRangeReturnsOnCall(i int, result1 error) {
	fake.DeleteIPRangeStub = nil
	if fake.deleteIPRangeReturnsOnCall == nil {
		fake.deleteIPRangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteIPRangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteTerminal(tx db.Transaction, terminalID int64) error {
	fake.deleteTerminalMutex.Lock()
	ret, specificReturn := fake.deleteTerminalReturnsOnCall[len(fake.deleteTerminalArgsForCall)]
	fake.deleteTerminalArgsForCall = append(fake.deleteTerminalArgsForCall, struct {
		tx         db.Transaction
		terminalID int64
	}{tx, terminalID})
	fake.recordInvocation("DeleteTerminal", []interface{}{tx, terminalID})
	fake.deleteTerminalMutex.Unlock()
	if fake.DeleteTerminalStub != nil {
		return fake.DeleteTerminalStub(tx, terminalID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteTerminalReturns.result1
}

func (fake *EgressPolicyRepo) DeleteTerminalCallCount() int {
	fake.deleteTerminalMutex.RLock()
	defer fake.deleteTerminalMutex.RUnlock()
	return len(fake.deleteTerminalArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteTerminalArgsForCall(i int) (db.Transaction, int64) {
	fake.deleteTerminalMutex.RLock()
	defer fake.deleteTerminalMutex.RUnlock()
	return fake.deleteTerminalArgsForCall[i].tx, fake.deleteTerminalArgsForCall[i].terminalID
}

func (fake *EgressPolicyRepo) DeleteTerminalReturns(result1 error) {
	fake.DeleteTerminalStub = nil
	fake.deleteTerminalReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteTerminalReturnsOnCall(i int, result1 error) {
	fake.DeleteTerminalStub = nil
	if fake.deleteTerminalReturnsOnCall == nil {
		fake.deleteTerminalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteTerminalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteApp(tx db.Transaction, appID int64) error {
	fake.deleteAppMutex.Lock()
	ret, specificReturn := fake.deleteAppReturnsOnCall[len(fake.deleteAppArgsForCall)]
	fake.deleteAppArgsForCall = append(fake.deleteAppArgsForCall, struct {
		tx    db.Transaction
		appID int64
	}{tx, appID})
	fake.recordInvocation("DeleteApp", []interface{}{tx, appID})
	fake.deleteAppMutex.Unlock()
	if fake.DeleteAppStub != nil {
		return fake.DeleteAppStub(tx, appID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteAppReturns.result1
}

func (fake *EgressPolicyRepo) DeleteAppCallCount() int {
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	return len(fake.deleteAppArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteAppArgsForCall(i int) (db.Transaction, int64) {
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	return fake.deleteAppArgsForCall[i].tx, fake.deleteAppArgsForCall[i].appID
}

func (fake *EgressPolicyRepo) DeleteAppReturns(result1 error) {
	fake.DeleteAppStub = nil
	fake.deleteAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteAppReturnsOnCall(i int, result1 error) {
	fake.DeleteAppStub = nil
	if fake.deleteAppReturnsOnCall == nil {
		fake.deleteAppReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAppReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) IsTerminalInUse(tx db.Transaction, terminalID int64) (bool, error) {
	fake.isTerminalInUseMutex.Lock()
	ret, specificReturn := fake.isTerminalInUseReturnsOnCall[len(fake.isTerminalInUseArgsForCall)]
	fake.isTerminalInUseArgsForCall = append(fake.isTerminalInUseArgsForCall, struct {
		tx         db.Transaction
		terminalID int64
	}{tx, terminalID})
	fake.recordInvocation("IsTerminalInUse", []interface{}{tx, terminalID})
	fake.isTerminalInUseMutex.Unlock()
	if fake.IsTerminalInUseStub != nil {
		return fake.IsTerminalInUseStub(tx, terminalID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isTerminalInUseReturns.result1, fake.isTerminalInUseReturns.result2
}

func (fake *EgressPolicyRepo) IsTerminalInUseCallCount() int {
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	return len(fake.isTerminalInUseArgsForCall)
}

func (fake *EgressPolicyRepo) IsTerminalInUseArgsForCall(i int) (db.Transaction, int64) {
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	return fake.isTerminalInUseArgsForCall[i].tx, fake.isTerminalInUseArgsForCall[i].terminalID
}

func (fake *EgressPolicyRepo) IsTerminalInUseReturns(result1 bool, result2 error) {
	fake.IsTerminalInUseStub = nil
	fake.isTerminalInUseReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) IsTerminalInUseReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsTerminalInUseStub = nil
	if fake.isTerminalInUseReturnsOnCall == nil {
		fake.isTerminalInUseReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isTerminalInUseReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTerminalMutex.RLock()
	defer fake.createTerminalMutex.RUnlock()
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	fake.getIDsByEgressPolicyMutex.RLock()
	defer fake.getIDsByEgressPolicyMutex.RUnlock()
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	fake.deleteIPRangeMutex.RLock()
	defer fake.deleteIPRangeMutex.RUnlock()
	fake.deleteTerminalMutex.RLock()
	defer fake.deleteTerminalMutex.RUnlock()
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyRepo) recordInvocation(key string, args []interface{}) {
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
