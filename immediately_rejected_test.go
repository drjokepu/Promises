package promises

import "testing"

func TestImmediatelyRejected(t *testing.T) {
	initialValue := "hello"
	promise := Reject(initialValue)
	promise.Then(func(value PromiseResult) Promise {
		t.Error("promise was resolved")
		return Empty()
	}).Fail(func(reason PromiseResult) Promise {
		if reason.(string) != initialValue {
			t.Error("%v != %v", initialValue, reason)
		}
		return Empty()
	})
}
