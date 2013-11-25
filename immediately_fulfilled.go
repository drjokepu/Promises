package promises

type immediatelyFulfilled struct {
	fulfilled chan PromiseResult
	rejected  chan PromiseResult
}

func (promise *immediatelyFulfilled) onFulfilled() chan PromiseResult {
	return promise.fulfilled
}

func (promise *immediatelyFulfilled) onRejected() chan PromiseResult {
	return promise.rejected
}

func (promise *immediatelyFulfilled) Then(next func(PromiseResult) Promise) Promise {
	return then(promise, next)
}

func (promise *immediatelyFulfilled) Fail(next func(PromiseResult) Promise) Promise {
	return fail(promise, next)
}

func newImmediatelyFulfilled(value PromiseResult) *immediatelyFulfilled {
	promise := &immediatelyFulfilled{
		fulfilled: make(chan PromiseResult, 1),
		rejected:  make(chan PromiseResult, 1),
	}
	promise.fulfilled <- value
	return promise
}

func Fulfil(value PromiseResult) Promise {
	return newImmediatelyFulfilled(value)
}

func Empty() Promise {
	return Fulfil(nil)
}
