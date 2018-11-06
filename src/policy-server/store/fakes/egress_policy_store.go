// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/store"
	"sync"
)

type EgressPolicyStore struct {
	CreateStub        func([]store.EgressPolicy) ([]store.EgressPolicy, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []store.EgressPolicy
	}
	createReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	DeleteStub        func(guids ...string) ([]store.EgressPolicy, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		guids []string
	}
	deleteReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	AllStub        func() ([]store.EgressPolicy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []store.EgressPolicy
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetBySourceGuidsStub        func(srcGuids []string) ([]store.EgressPolicy, error)
	getBySourceGuidsMutex       sync.RWMutex
	getBySourceGuidsArgsForCall []struct {
		srcGuids []string
	}
	getBySourceGuidsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getBySourceGuidsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetByFilterStub        func(sourceIds, sourceTypes, destinationIds, destinationNames []string) ([]store.EgressPolicy, error)
	getByFilterMutex       sync.RWMutex
	getByFilterArgsForCall []struct {
		sourceIds        []string
		sourceTypes      []string
		destinationIds   []string
		destinationNames []string
	}
	getByFilterReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getByFilterReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyStore) Create(arg1 []store.EgressPolicy) ([]store.EgressPolicy, error) {
	var arg1Copy []store.EgressPolicy
	if arg1 != nil {
		arg1Copy = make([]store.EgressPolicy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []store.EgressPolicy
	}{arg1Copy})
	fake.recordInvocation("Create", []interface{}{arg1Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *EgressPolicyStore) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *EgressPolicyStore) CreateArgsForCall(i int) []store.EgressPolicy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *EgressPolicyStore) CreateReturns(result1 []store.EgressPolicy, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) CreateReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Delete(guids ...string) ([]store.EgressPolicy, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		guids []string
	}{guids})
	fake.recordInvocation("Delete", []interface{}{guids})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(guids...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteReturns.result1, fake.deleteReturns.result2
}

func (fake *EgressPolicyStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *EgressPolicyStore) DeleteArgsForCall(i int) []string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].guids
}

func (fake *EgressPolicyStore) DeleteReturns(result1 []store.EgressPolicy, result2 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) DeleteReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) All() ([]store.EgressPolicy, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *EgressPolicyStore) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *EgressPolicyStore) AllReturns(result1 []store.EgressPolicy, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) AllReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetBySourceGuids(srcGuids []string) ([]store.EgressPolicy, error) {
	var srcGuidsCopy []string
	if srcGuids != nil {
		srcGuidsCopy = make([]string, len(srcGuids))
		copy(srcGuidsCopy, srcGuids)
	}
	fake.getBySourceGuidsMutex.Lock()
	ret, specificReturn := fake.getBySourceGuidsReturnsOnCall[len(fake.getBySourceGuidsArgsForCall)]
	fake.getBySourceGuidsArgsForCall = append(fake.getBySourceGuidsArgsForCall, struct {
		srcGuids []string
	}{srcGuidsCopy})
	fake.recordInvocation("GetBySourceGuids", []interface{}{srcGuidsCopy})
	fake.getBySourceGuidsMutex.Unlock()
	if fake.GetBySourceGuidsStub != nil {
		return fake.GetBySourceGuidsStub(srcGuids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getBySourceGuidsReturns.result1, fake.getBySourceGuidsReturns.result2
}

func (fake *EgressPolicyStore) GetBySourceGuidsCallCount() int {
	fake.getBySourceGuidsMutex.RLock()
	defer fake.getBySourceGuidsMutex.RUnlock()
	return len(fake.getBySourceGuidsArgsForCall)
}

func (fake *EgressPolicyStore) GetBySourceGuidsArgsForCall(i int) []string {
	fake.getBySourceGuidsMutex.RLock()
	defer fake.getBySourceGuidsMutex.RUnlock()
	return fake.getBySourceGuidsArgsForCall[i].srcGuids
}

func (fake *EgressPolicyStore) GetBySourceGuidsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetBySourceGuidsStub = nil
	fake.getBySourceGuidsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetBySourceGuidsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetBySourceGuidsStub = nil
	if fake.getBySourceGuidsReturnsOnCall == nil {
		fake.getBySourceGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getBySourceGuidsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetByFilter(sourceIds []string, sourceTypes []string, destinationIds []string, destinationNames []string) ([]store.EgressPolicy, error) {
	var sourceIdsCopy []string
	if sourceIds != nil {
		sourceIdsCopy = make([]string, len(sourceIds))
		copy(sourceIdsCopy, sourceIds)
	}
	var sourceTypesCopy []string
	if sourceTypes != nil {
		sourceTypesCopy = make([]string, len(sourceTypes))
		copy(sourceTypesCopy, sourceTypes)
	}
	var destinationIdsCopy []string
	if destinationIds != nil {
		destinationIdsCopy = make([]string, len(destinationIds))
		copy(destinationIdsCopy, destinationIds)
	}
	var destinationNamesCopy []string
	if destinationNames != nil {
		destinationNamesCopy = make([]string, len(destinationNames))
		copy(destinationNamesCopy, destinationNames)
	}
	fake.getByFilterMutex.Lock()
	ret, specificReturn := fake.getByFilterReturnsOnCall[len(fake.getByFilterArgsForCall)]
	fake.getByFilterArgsForCall = append(fake.getByFilterArgsForCall, struct {
		sourceIds        []string
		sourceTypes      []string
		destinationIds   []string
		destinationNames []string
	}{sourceIdsCopy, sourceTypesCopy, destinationIdsCopy, destinationNamesCopy})
	fake.recordInvocation("GetByFilter", []interface{}{sourceIdsCopy, sourceTypesCopy, destinationIdsCopy, destinationNamesCopy})
	fake.getByFilterMutex.Unlock()
	if fake.GetByFilterStub != nil {
		return fake.GetByFilterStub(sourceIds, sourceTypes, destinationIds, destinationNames)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByFilterReturns.result1, fake.getByFilterReturns.result2
}

func (fake *EgressPolicyStore) GetByFilterCallCount() int {
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	return len(fake.getByFilterArgsForCall)
}

func (fake *EgressPolicyStore) GetByFilterArgsForCall(i int) ([]string, []string, []string, []string) {
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	return fake.getByFilterArgsForCall[i].sourceIds, fake.getByFilterArgsForCall[i].sourceTypes, fake.getByFilterArgsForCall[i].destinationIds, fake.getByFilterArgsForCall[i].destinationNames
}

func (fake *EgressPolicyStore) GetByFilterReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetByFilterStub = nil
	fake.getByFilterReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) GetByFilterReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetByFilterStub = nil
	if fake.getByFilterReturnsOnCall == nil {
		fake.getByFilterReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getByFilterReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.getBySourceGuidsMutex.RLock()
	defer fake.getBySourceGuidsMutex.RUnlock()
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	return fake.invocations
}

func (fake *EgressPolicyStore) recordInvocation(key string, args []interface{}) {
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
