package promises

import "sync"

type Deferred struct {
	stateOnce *sync.Once
	fulfilled chan PromiseResult
	rejected  chan PromiseResult
}

func (promise *Deferred) onFulfilled() chan PromiseResult {
	return promise.fulfilled
}

func (promise *Deferred) onRejected() chan PromiseResult {
	return promise.rejected
}

func (promise *Deferred) Then(next func(PromiseResult) Promise) Promise {
	return then(promise, next)
}

func (promise *Deferred) Fail(next func(PromiseResult) Promise) Promise {
	return fail(promise, next)
}

func (promise *Deferred) Resolve(value PromiseResult) {
	promise.stateOnce.Do(func() {
		promise.fulfilled <- value
	})
}

func (promise *Deferred) Reject(reason PromiseResult) {
	promise.stateOnce.Do(func() {
		promise.rejected <- reason
	})
}

func newDeferred() *Deferred {
	promise := &Deferred{
		stateOnce: new(sync.Once),
		fulfilled: make(chan PromiseResult, 1),
		rejected:  make(chan PromiseResult, 1),
	}
	return promise
}

func Defer() *Deferred {
	return newDeferred()
}
