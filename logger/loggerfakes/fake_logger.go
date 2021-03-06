// This file was generated by counterfeiter
package loggerfakes

import (
	"sync"

	"github.com/robdimsdale/concourse-pipeline-resource/logger"
)

type FakeLogger struct {
	DebugfStub        func(format string, a ...interface{}) (n int, err error)
	debugfMutex       sync.RWMutex
	debugfArgsForCall []struct {
		format string
		a      []interface{}
	}
	debugfReturns struct {
		result1 int
		result2 error
	}
}

func (fake *FakeLogger) Debugf(format string, a ...interface{}) (n int, err error) {
	fake.debugfMutex.Lock()
	fake.debugfArgsForCall = append(fake.debugfArgsForCall, struct {
		format string
		a      []interface{}
	}{format, a})
	fake.debugfMutex.Unlock()
	if fake.DebugfStub != nil {
		return fake.DebugfStub(format, a...)
	} else {
		return fake.debugfReturns.result1, fake.debugfReturns.result2
	}
}

func (fake *FakeLogger) DebugfCallCount() int {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return len(fake.debugfArgsForCall)
}

func (fake *FakeLogger) DebugfArgsForCall(i int) (string, []interface{}) {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return fake.debugfArgsForCall[i].format, fake.debugfArgsForCall[i].a
}

func (fake *FakeLogger) DebugfReturns(result1 int, result2 error) {
	fake.DebugfStub = nil
	fake.debugfReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ logger.Logger = new(FakeLogger)
