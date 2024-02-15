// Code generated by counterfeiter. DO NOT EDIT.
package telemetryfakes

import (
	"sync"

	"github.com/nginxinc/nginx-gateway-fabric/internal/mode/static/state/graph"
	"github.com/nginxinc/nginx-gateway-fabric/internal/mode/static/telemetry"
)

type FakeGraphGetter struct {
	GetLatestGraphStub        func() *graph.Graph
	getLatestGraphMutex       sync.RWMutex
	getLatestGraphArgsForCall []struct {
	}
	getLatestGraphReturns struct {
		result1 *graph.Graph
	}
	getLatestGraphReturnsOnCall map[int]struct {
		result1 *graph.Graph
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGraphGetter) GetLatestGraph() *graph.Graph {
	fake.getLatestGraphMutex.Lock()
	ret, specificReturn := fake.getLatestGraphReturnsOnCall[len(fake.getLatestGraphArgsForCall)]
	fake.getLatestGraphArgsForCall = append(fake.getLatestGraphArgsForCall, struct {
	}{})
	stub := fake.GetLatestGraphStub
	fakeReturns := fake.getLatestGraphReturns
	fake.recordInvocation("GetLatestGraph", []interface{}{})
	fake.getLatestGraphMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGraphGetter) GetLatestGraphCallCount() int {
	fake.getLatestGraphMutex.RLock()
	defer fake.getLatestGraphMutex.RUnlock()
	return len(fake.getLatestGraphArgsForCall)
}

func (fake *FakeGraphGetter) GetLatestGraphCalls(stub func() *graph.Graph) {
	fake.getLatestGraphMutex.Lock()
	defer fake.getLatestGraphMutex.Unlock()
	fake.GetLatestGraphStub = stub
}

func (fake *FakeGraphGetter) GetLatestGraphReturns(result1 *graph.Graph) {
	fake.getLatestGraphMutex.Lock()
	defer fake.getLatestGraphMutex.Unlock()
	fake.GetLatestGraphStub = nil
	fake.getLatestGraphReturns = struct {
		result1 *graph.Graph
	}{result1}
}

func (fake *FakeGraphGetter) GetLatestGraphReturnsOnCall(i int, result1 *graph.Graph) {
	fake.getLatestGraphMutex.Lock()
	defer fake.getLatestGraphMutex.Unlock()
	fake.GetLatestGraphStub = nil
	if fake.getLatestGraphReturnsOnCall == nil {
		fake.getLatestGraphReturnsOnCall = make(map[int]struct {
			result1 *graph.Graph
		})
	}
	fake.getLatestGraphReturnsOnCall[i] = struct {
		result1 *graph.Graph
	}{result1}
}

func (fake *FakeGraphGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getLatestGraphMutex.RLock()
	defer fake.getLatestGraphMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGraphGetter) recordInvocation(key string, args []interface{}) {
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

var _ telemetry.GraphGetter = new(FakeGraphGetter)