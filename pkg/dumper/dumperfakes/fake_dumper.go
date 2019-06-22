// Code generated by counterfeiter. DO NOT EDIT.
package dumperfakes

import (
	"context"
	"io"
	"sync"

	"github.com/bipol/scrapedumper/pkg/dumper"
)

type FakeDumper struct {
	DumpStub        func(context.Context, io.Reader, string) error
	dumpMutex       sync.RWMutex
	dumpArgsForCall []struct {
		arg1 context.Context
		arg2 io.Reader
		arg3 string
	}
	dumpReturns struct {
		result1 error
	}
	dumpReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDumper) Dump(arg1 context.Context, arg2 io.Reader, arg3 string) error {
	fake.dumpMutex.Lock()
	ret, specificReturn := fake.dumpReturnsOnCall[len(fake.dumpArgsForCall)]
	fake.dumpArgsForCall = append(fake.dumpArgsForCall, struct {
		arg1 context.Context
		arg2 io.Reader
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Dump", []interface{}{arg1, arg2, arg3})
	fake.dumpMutex.Unlock()
	if fake.DumpStub != nil {
		return fake.DumpStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dumpReturns
	return fakeReturns.result1
}

func (fake *FakeDumper) DumpCallCount() int {
	fake.dumpMutex.RLock()
	defer fake.dumpMutex.RUnlock()
	return len(fake.dumpArgsForCall)
}

func (fake *FakeDumper) DumpCalls(stub func(context.Context, io.Reader, string) error) {
	fake.dumpMutex.Lock()
	defer fake.dumpMutex.Unlock()
	fake.DumpStub = stub
}

func (fake *FakeDumper) DumpArgsForCall(i int) (context.Context, io.Reader, string) {
	fake.dumpMutex.RLock()
	defer fake.dumpMutex.RUnlock()
	argsForCall := fake.dumpArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDumper) DumpReturns(result1 error) {
	fake.dumpMutex.Lock()
	defer fake.dumpMutex.Unlock()
	fake.DumpStub = nil
	fake.dumpReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDumper) DumpReturnsOnCall(i int, result1 error) {
	fake.dumpMutex.Lock()
	defer fake.dumpMutex.Unlock()
	fake.DumpStub = nil
	if fake.dumpReturnsOnCall == nil {
		fake.dumpReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.dumpReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDumper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dumpMutex.RLock()
	defer fake.dumpMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDumper) recordInvocation(key string, args []interface{}) {
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

var _ dumper.Dumper = new(FakeDumper)
