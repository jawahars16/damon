// Code generated by counterfeiter. DO NOT EDIT.
package nomadfakes

import (
	"sync"

	"github.com/hashicorp/nomad/api"
	"github.com/hcjulz/damon/nomad"
)

type FakeAllocationsClient struct {
	ListStub        func(*api.QueryOptions) ([]*api.AllocationListStub, *api.QueryMeta, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 *api.QueryOptions
	}
	listReturns struct {
		result1 []*api.AllocationListStub
		result2 *api.QueryMeta
		result3 error
	}
	listReturnsOnCall map[int]struct {
		result1 []*api.AllocationListStub
		result2 *api.QueryMeta
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAllocationsClient) List(arg1 *api.QueryOptions) ([]*api.AllocationListStub, *api.QueryMeta, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 *api.QueryOptions
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAllocationsClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeAllocationsClient) ListCalls(stub func(*api.QueryOptions) ([]*api.AllocationListStub, *api.QueryMeta, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeAllocationsClient) ListArgsForCall(i int) *api.QueryOptions {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAllocationsClient) ListReturns(result1 []*api.AllocationListStub, result2 *api.QueryMeta, result3 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*api.AllocationListStub
		result2 *api.QueryMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAllocationsClient) ListReturnsOnCall(i int, result1 []*api.AllocationListStub, result2 *api.QueryMeta, result3 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []*api.AllocationListStub
			result2 *api.QueryMeta
			result3 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []*api.AllocationListStub
		result2 *api.QueryMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAllocationsClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAllocationsClient) recordInvocation(key string, args []interface{}) {
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

var _ nomad.AllocationsClient = new(FakeAllocationsClient)