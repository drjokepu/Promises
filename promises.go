package promises

const (
	Pending = iota
	Fulfilled
	Rejected
)

type PromiseState int

type PromiseResult interface {
}

type Promise interface {
	onFulfilled() chan PromiseResult
	onRejected() chan PromiseResult
	Then(func(PromiseResult) Promise) Promise
	Fail(func(PromiseResult) Promise) Promise
}

func thenOrFail(promise Promise, thenCallback func(PromiseResult) Promise, failCallback func(PromiseResult) Promise) {
	
}

func then(promise Promise, callback func(PromiseResult) Promise) Promise {
	deferred := Defer()
	go func() {
		select {
		case value := <-promise.onFulfilled():
			deferred.Resolve(callback(value))
			return
		case reason := <-promise.onRejected():
			deferred.Reject(newImmediatelyRejected(reason))
		}
	}()
	return deferred
}

func fail(promise Promise, callback func(PromiseResult) Promise) Promise {
	deferred := Defer()
	go func() {
		select {
		case value := <-promise.onFulfilled():
			deferred.Resolve(newImmediatelyFulfilled(value))
			return
		case reason := <-promise.onRejected():
			deferred.Resolve(callback(reason))
			return
		}
	}()
	return deferred
}

func Do(action func() Promise) Promise {
	deferred := Defer()
	go func() {
		innerPromise := action()
		select {
		case value := <-innerPromise.onFulfilled():
			deferred.Resolve(value)
			return
		case reason := <-innerPromise.onRejected():
			deferred.Reject(reason)
			return
		}
	}()
	return deferred
}
