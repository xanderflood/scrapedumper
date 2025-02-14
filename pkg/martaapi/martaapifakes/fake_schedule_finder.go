// Code generated by counterfeiter. DO NOT EDIT.
package martaapifakes

import (
	"context"
	"io"
	"sync"

	"github.com/bipol/scrapedumper/pkg/martaapi"
)

type FakeScheduleFinder struct {
	FindSchedulesStub        func(context.Context) (io.ReadCloser, error)
	findSchedulesMutex       sync.RWMutex
	findSchedulesArgsForCall []struct {
		arg1 context.Context
	}
	findSchedulesReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	findSchedulesReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScheduleFinder) FindSchedules(arg1 context.Context) (io.ReadCloser, error) {
	fake.findSchedulesMutex.Lock()
	ret, specificReturn := fake.findSchedulesReturnsOnCall[len(fake.findSchedulesArgsForCall)]
	fake.findSchedulesArgsForCall = append(fake.findSchedulesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("FindSchedules", []interface{}{arg1})
	fake.findSchedulesMutex.Unlock()
	if fake.FindSchedulesStub != nil {
		return fake.FindSchedulesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findSchedulesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeScheduleFinder) FindSchedulesCallCount() int {
	fake.findSchedulesMutex.RLock()
	defer fake.findSchedulesMutex.RUnlock()
	return len(fake.findSchedulesArgsForCall)
}

func (fake *FakeScheduleFinder) FindSchedulesCalls(stub func(context.Context) (io.ReadCloser, error)) {
	fake.findSchedulesMutex.Lock()
	defer fake.findSchedulesMutex.Unlock()
	fake.FindSchedulesStub = stub
}

func (fake *FakeScheduleFinder) FindSchedulesArgsForCall(i int) context.Context {
	fake.findSchedulesMutex.RLock()
	defer fake.findSchedulesMutex.RUnlock()
	argsForCall := fake.findSchedulesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeScheduleFinder) FindSchedulesReturns(result1 io.ReadCloser, result2 error) {
	fake.findSchedulesMutex.Lock()
	defer fake.findSchedulesMutex.Unlock()
	fake.FindSchedulesStub = nil
	fake.findSchedulesReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeScheduleFinder) FindSchedulesReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.findSchedulesMutex.Lock()
	defer fake.findSchedulesMutex.Unlock()
	fake.FindSchedulesStub = nil
	if fake.findSchedulesReturnsOnCall == nil {
		fake.findSchedulesReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.findSchedulesReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeScheduleFinder) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.typeReturns
	return fakeReturns.result1
}

func (fake *FakeScheduleFinder) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeScheduleFinder) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeScheduleFinder) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeScheduleFinder) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeScheduleFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findSchedulesMutex.RLock()
	defer fake.findSchedulesMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScheduleFinder) recordInvocation(key string, args []interface{}) {
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

var _ martaapi.ScheduleFinder = new(FakeScheduleFinder)
