// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/api"
	"sync"
)

type CCClient struct {
	GetAppSpacesStub        func(token string, appGUIDs []string) (map[string]string, error)
	getAppSpacesMutex       sync.RWMutex
	getAppSpacesArgsForCall []struct {
		token    string
		appGUIDs []string
	}
	getAppSpacesReturns struct {
		result1 map[string]string
		result2 error
	}
	getAppSpacesReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetSpaceStub        func(token, spaceGUID string) (*api.Space, error)
	getSpaceMutex       sync.RWMutex
	getSpaceArgsForCall []struct {
		token     string
		spaceGUID string
	}
	getSpaceReturns struct {
		result1 *api.Space
		result2 error
	}
	getSpaceReturnsOnCall map[int]struct {
		result1 *api.Space
		result2 error
	}
	GetSpaceGUIDsStub        func(token string, appGUIDs []string) ([]string, error)
	getSpaceGUIDsMutex       sync.RWMutex
	getSpaceGUIDsArgsForCall []struct {
		token    string
		appGUIDs []string
	}
	getSpaceGUIDsReturns struct {
		result1 []string
		result2 error
	}
	getSpaceGUIDsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetSubjectSpaceStub        func(token, subjectId string, spaces api.Space) (*api.Space, error)
	getSubjectSpaceMutex       sync.RWMutex
	getSubjectSpaceArgsForCall []struct {
		token     string
		subjectId string
		spaces    api.Space
	}
	getSubjectSpaceReturns struct {
		result1 *api.Space
		result2 error
	}
	getSubjectSpaceReturnsOnCall map[int]struct {
		result1 *api.Space
		result2 error
	}
	GetSubjectSpacesStub        func(token, subjectId string) (map[string]struct{}, error)
	getSubjectSpacesMutex       sync.RWMutex
	getSubjectSpacesArgsForCall []struct {
		token     string
		subjectId string
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

func (fake *CCClient) GetAppSpaces(token string, appGUIDs []string) (map[string]string, error) {
	var appGUIDsCopy []string
	if appGUIDs != nil {
		appGUIDsCopy = make([]string, len(appGUIDs))
		copy(appGUIDsCopy, appGUIDs)
	}
	fake.getAppSpacesMutex.Lock()
	ret, specificReturn := fake.getAppSpacesReturnsOnCall[len(fake.getAppSpacesArgsForCall)]
	fake.getAppSpacesArgsForCall = append(fake.getAppSpacesArgsForCall, struct {
		token    string
		appGUIDs []string
	}{token, appGUIDsCopy})
	fake.recordInvocation("GetAppSpaces", []interface{}{token, appGUIDsCopy})
	fake.getAppSpacesMutex.Unlock()
	if fake.GetAppSpacesStub != nil {
		return fake.GetAppSpacesStub(token, appGUIDs)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAppSpacesReturns.result1, fake.getAppSpacesReturns.result2
}

func (fake *CCClient) GetAppSpacesCallCount() int {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	return len(fake.getAppSpacesArgsForCall)
}

func (fake *CCClient) GetAppSpacesArgsForCall(i int) (string, []string) {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	return fake.getAppSpacesArgsForCall[i].token, fake.getAppSpacesArgsForCall[i].appGUIDs
}

func (fake *CCClient) GetAppSpacesReturns(result1 map[string]string, result2 error) {
	fake.GetAppSpacesStub = nil
	fake.getAppSpacesReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetAppSpacesReturnsOnCall(i int, result1 map[string]string, result2 error) {
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

func (fake *CCClient) GetSpace(token string, spaceGUID string) (*api.Space, error) {
	fake.getSpaceMutex.Lock()
	ret, specificReturn := fake.getSpaceReturnsOnCall[len(fake.getSpaceArgsForCall)]
	fake.getSpaceArgsForCall = append(fake.getSpaceArgsForCall, struct {
		token     string
		spaceGUID string
	}{token, spaceGUID})
	fake.recordInvocation("GetSpace", []interface{}{token, spaceGUID})
	fake.getSpaceMutex.Unlock()
	if fake.GetSpaceStub != nil {
		return fake.GetSpaceStub(token, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceReturns.result1, fake.getSpaceReturns.result2
}

func (fake *CCClient) GetSpaceCallCount() int {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return len(fake.getSpaceArgsForCall)
}

func (fake *CCClient) GetSpaceArgsForCall(i int) (string, string) {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return fake.getSpaceArgsForCall[i].token, fake.getSpaceArgsForCall[i].spaceGUID
}

func (fake *CCClient) GetSpaceReturns(result1 *api.Space, result2 error) {
	fake.GetSpaceStub = nil
	fake.getSpaceReturns = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceReturnsOnCall(i int, result1 *api.Space, result2 error) {
	fake.GetSpaceStub = nil
	if fake.getSpaceReturnsOnCall == nil {
		fake.getSpaceReturnsOnCall = make(map[int]struct {
			result1 *api.Space
			result2 error
		})
	}
	fake.getSpaceReturnsOnCall[i] = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceGUIDs(token string, appGUIDs []string) ([]string, error) {
	var appGUIDsCopy []string
	if appGUIDs != nil {
		appGUIDsCopy = make([]string, len(appGUIDs))
		copy(appGUIDsCopy, appGUIDs)
	}
	fake.getSpaceGUIDsMutex.Lock()
	ret, specificReturn := fake.getSpaceGUIDsReturnsOnCall[len(fake.getSpaceGUIDsArgsForCall)]
	fake.getSpaceGUIDsArgsForCall = append(fake.getSpaceGUIDsArgsForCall, struct {
		token    string
		appGUIDs []string
	}{token, appGUIDsCopy})
	fake.recordInvocation("GetSpaceGUIDs", []interface{}{token, appGUIDsCopy})
	fake.getSpaceGUIDsMutex.Unlock()
	if fake.GetSpaceGUIDsStub != nil {
		return fake.GetSpaceGUIDsStub(token, appGUIDs)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceGUIDsReturns.result1, fake.getSpaceGUIDsReturns.result2
}

func (fake *CCClient) GetSpaceGUIDsCallCount() int {
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	return len(fake.getSpaceGUIDsArgsForCall)
}

func (fake *CCClient) GetSpaceGUIDsArgsForCall(i int) (string, []string) {
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	return fake.getSpaceGUIDsArgsForCall[i].token, fake.getSpaceGUIDsArgsForCall[i].appGUIDs
}

func (fake *CCClient) GetSpaceGUIDsReturns(result1 []string, result2 error) {
	fake.GetSpaceGUIDsStub = nil
	fake.getSpaceGUIDsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceGUIDsReturnsOnCall(i int, result1 []string, result2 error) {
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

func (fake *CCClient) GetSubjectSpace(token string, subjectId string, spaces api.Space) (*api.Space, error) {
	fake.getSubjectSpaceMutex.Lock()
	ret, specificReturn := fake.getSubjectSpaceReturnsOnCall[len(fake.getSubjectSpaceArgsForCall)]
	fake.getSubjectSpaceArgsForCall = append(fake.getSubjectSpaceArgsForCall, struct {
		token     string
		subjectId string
		spaces    api.Space
	}{token, subjectId, spaces})
	fake.recordInvocation("GetSubjectSpace", []interface{}{token, subjectId, spaces})
	fake.getSubjectSpaceMutex.Unlock()
	if fake.GetSubjectSpaceStub != nil {
		return fake.GetSubjectSpaceStub(token, subjectId, spaces)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSubjectSpaceReturns.result1, fake.getSubjectSpaceReturns.result2
}

func (fake *CCClient) GetSubjectSpaceCallCount() int {
	fake.getSubjectSpaceMutex.RLock()
	defer fake.getSubjectSpaceMutex.RUnlock()
	return len(fake.getSubjectSpaceArgsForCall)
}

func (fake *CCClient) GetSubjectSpaceArgsForCall(i int) (string, string, api.Space) {
	fake.getSubjectSpaceMutex.RLock()
	defer fake.getSubjectSpaceMutex.RUnlock()
	return fake.getSubjectSpaceArgsForCall[i].token, fake.getSubjectSpaceArgsForCall[i].subjectId, fake.getSubjectSpaceArgsForCall[i].spaces
}

func (fake *CCClient) GetSubjectSpaceReturns(result1 *api.Space, result2 error) {
	fake.GetSubjectSpaceStub = nil
	fake.getSubjectSpaceReturns = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSubjectSpaceReturnsOnCall(i int, result1 *api.Space, result2 error) {
	fake.GetSubjectSpaceStub = nil
	if fake.getSubjectSpaceReturnsOnCall == nil {
		fake.getSubjectSpaceReturnsOnCall = make(map[int]struct {
			result1 *api.Space
			result2 error
		})
	}
	fake.getSubjectSpaceReturnsOnCall[i] = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSubjectSpaces(token string, subjectId string) (map[string]struct{}, error) {
	fake.getSubjectSpacesMutex.Lock()
	ret, specificReturn := fake.getSubjectSpacesReturnsOnCall[len(fake.getSubjectSpacesArgsForCall)]
	fake.getSubjectSpacesArgsForCall = append(fake.getSubjectSpacesArgsForCall, struct {
		token     string
		subjectId string
	}{token, subjectId})
	fake.recordInvocation("GetSubjectSpaces", []interface{}{token, subjectId})
	fake.getSubjectSpacesMutex.Unlock()
	if fake.GetSubjectSpacesStub != nil {
		return fake.GetSubjectSpacesStub(token, subjectId)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSubjectSpacesReturns.result1, fake.getSubjectSpacesReturns.result2
}

func (fake *CCClient) GetSubjectSpacesCallCount() int {
	fake.getSubjectSpacesMutex.RLock()
	defer fake.getSubjectSpacesMutex.RUnlock()
	return len(fake.getSubjectSpacesArgsForCall)
}

func (fake *CCClient) GetSubjectSpacesArgsForCall(i int) (string, string) {
	fake.getSubjectSpacesMutex.RLock()
	defer fake.getSubjectSpacesMutex.RUnlock()
	return fake.getSubjectSpacesArgsForCall[i].token, fake.getSubjectSpacesArgsForCall[i].subjectId
}

func (fake *CCClient) GetSubjectSpacesReturns(result1 map[string]struct{}, result2 error) {
	fake.GetSubjectSpacesStub = nil
	fake.getSubjectSpacesReturns = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSubjectSpacesReturnsOnCall(i int, result1 map[string]struct{}, result2 error) {
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
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	fake.getSubjectSpaceMutex.RLock()
	defer fake.getSubjectSpaceMutex.RUnlock()
	fake.getSubjectSpacesMutex.RLock()
	defer fake.getSubjectSpacesMutex.RUnlock()
	return fake.invocations
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
