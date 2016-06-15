// This file was generated by counterfeiter
package netfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/net"
)

type FakeReadSeekCloser struct {
	ReadStub        func(p []byte) (n int, err error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		p []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	SeekStub        func(offset int64, whence int) (int64, error)
	seekMutex       sync.RWMutex
	seekArgsForCall []struct {
		offset int64
		whence int
	}
	seekReturns struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReadSeekCloser) Read(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.recordInvocation("Read", []interface{}{pCopy})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(p)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeReadSeekCloser) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeReadSeekCloser) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].p
}

func (fake *FakeReadSeekCloser) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeReadSeekCloser) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeReadSeekCloser) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeReadSeekCloser) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReadSeekCloser) Seek(offset int64, whence int) (int64, error) {
	fake.seekMutex.Lock()
	fake.seekArgsForCall = append(fake.seekArgsForCall, struct {
		offset int64
		whence int
	}{offset, whence})
	fake.recordInvocation("Seek", []interface{}{offset, whence})
	fake.seekMutex.Unlock()
	if fake.SeekStub != nil {
		return fake.SeekStub(offset, whence)
	} else {
		return fake.seekReturns.result1, fake.seekReturns.result2
	}
}

func (fake *FakeReadSeekCloser) SeekCallCount() int {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return len(fake.seekArgsForCall)
}

func (fake *FakeReadSeekCloser) SeekArgsForCall(i int) (int64, int) {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return fake.seekArgsForCall[i].offset, fake.seekArgsForCall[i].whence
}

func (fake *FakeReadSeekCloser) SeekReturns(result1 int64, result2 error) {
	fake.SeekStub = nil
	fake.seekReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeReadSeekCloser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeReadSeekCloser) recordInvocation(key string, args []interface{}) {
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

var _ net.ReadSeekCloser = new(FakeReadSeekCloser)
