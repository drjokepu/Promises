package promises

func done(promise Promise) {
	fail(promise, func(reason PromiseResult) Promise {
		panic(reason)
	})
}
