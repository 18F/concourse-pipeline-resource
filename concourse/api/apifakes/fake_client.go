// This file was generated by counterfeiter
package apifakes

import (
	"sync"

	"github.com/robdimsdale/concourse-pipeline-resource/concourse/api"
)

type FakeClient struct {
	PipelinesStub        func() ([]api.Pipeline, error)
	pipelinesMutex       sync.RWMutex
	pipelinesArgsForCall []struct{}
	pipelinesReturns     struct {
		result1 []api.Pipeline
		result2 error
	}
}

func (fake *FakeClient) Pipelines() ([]api.Pipeline, error) {
	fake.pipelinesMutex.Lock()
	fake.pipelinesArgsForCall = append(fake.pipelinesArgsForCall, struct{}{})
	fake.pipelinesMutex.Unlock()
	if fake.PipelinesStub != nil {
		return fake.PipelinesStub()
	} else {
		return fake.pipelinesReturns.result1, fake.pipelinesReturns.result2
	}
}

func (fake *FakeClient) PipelinesCallCount() int {
	fake.pipelinesMutex.RLock()
	defer fake.pipelinesMutex.RUnlock()
	return len(fake.pipelinesArgsForCall)
}

func (fake *FakeClient) PipelinesReturns(result1 []api.Pipeline, result2 error) {
	fake.PipelinesStub = nil
	fake.pipelinesReturns = struct {
		result1 []api.Pipeline
		result2 error
	}{result1, result2}
}

var _ api.Client = new(FakeClient)
