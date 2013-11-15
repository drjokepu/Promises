package promises

import "testing"

func TestDeferResolved(t *testing.T) {
	initialValue := 500
	promise := Defer()
	defer func() {
		promise.Resolve(initialValue)
	}()

	promise.Then(func(value PromiseResult) Promise {
		if value.(int) != initialValue {
			t.Error("%v != %v", initialValue, value)
		}
		return Empty()
	}).Fail(func(value PromiseResult) Promise {
		t.Error("promise was rejected")
		return Empty()
	})
}
