package promises

import "testing"

func TestImmediatelyFulfilled(t *testing.T) {
	initialValue := 500
	promise := Fulfil(initialValue)
	promise.Then(func(value PromiseResult) Promise {
		if value.(int) != initialValue {
			t.Error("%v != %v", initialValue, value)
		}
		return Empty()
	}).Fail(func(reason PromiseResult) Promise {
		t.Error("promise was rejected")
		return Empty()
	})
}
