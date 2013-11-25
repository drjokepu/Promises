package promises

type immediatelyRejected struct {
	fulfilled chan PromiseResult
	rejected  chan PromiseResult
}

func (promise *immediatelyRejected) onFulfilled() chan PromiseResult {
	return promise.fulfilled
}

func (promise *immediatelyRejected) onRejected() chan PromiseResult {
	return promise.rejected
}

func (promise *immediatelyRejected) Then(next func(PromiseResult) Promise) Promise {
	return then(promise, next)
}

func (promise *immediatelyRejected) Fail(next func(PromiseResult) Promise) Promise {
	return fail(promise, next)
}

func newImmediatelyRejected(reason PromiseResult) *immediatelyRejected {
	promise := &immediatelyRejected{
		fulfilled: make(chan PromiseResult, 1),
		rejected:  make(chan PromiseResult, 1),
	}
	promise.rejected <- reason
	return promise
}

func Reject(reason PromiseResult) Promise {
	return newImmediatelyRejected(reason)
}
