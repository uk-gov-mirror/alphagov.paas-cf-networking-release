// This file was generated by counterfeiter
package fakes

import (
	"policy-server/models"
	"policy-server/store"
	"sync"
)

type Store struct {
	CreateStub        func([]models.Policy) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []models.Policy
	}
	createReturns struct {
		result1 error
	}
	AllStub        func() ([]models.Policy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []models.Policy
		result2 error
	}
	DeleteStub        func([]models.Policy) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []models.Policy
	}
	deleteReturns struct {
		result1 error
	}
	TagsStub        func() ([]models.Tag, error)
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct{}
	tagsReturns     struct {
		result1 []models.Tag
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Store) Create(arg1 []models.Policy) error {
	var arg1Copy []models.Policy
	if arg1 != nil {
		arg1Copy = make([]models.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []models.Policy
	}{arg1Copy})
	fake.recordInvocation("Create", []interface{}{arg1Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	return fake.createReturns.result1
}

func (fake *Store) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Store) CreateArgsForCall(i int) []models.Policy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *Store) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) All() ([]models.Policy, error) {
	fake.allMutex.Lock()
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *Store) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *Store) AllReturns(result1 []models.Policy, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []models.Policy
		result2 error
	}{result1, result2}
}

func (fake *Store) Delete(arg1 []models.Policy) error {
	var arg1Copy []models.Policy
	if arg1 != nil {
		arg1Copy = make([]models.Policy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []models.Policy
	}{arg1Copy})
	fake.recordInvocation("Delete", []interface{}{arg1Copy})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	return fake.deleteReturns.result1
}

func (fake *Store) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *Store) DeleteArgsForCall(i int) []models.Policy {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *Store) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) Tags() ([]models.Tag, error) {
	fake.tagsMutex.Lock()
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct{}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	return fake.tagsReturns.result1, fake.tagsReturns.result2
}

func (fake *Store) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *Store) TagsReturns(result1 []models.Tag, result2 error) {
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 []models.Tag
		result2 error
	}{result1, result2}
}

func (fake *Store) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return fake.invocations
}

func (fake *Store) recordInvocation(key string, args []interface{}) {
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

var _ store.Store = new(Store)
