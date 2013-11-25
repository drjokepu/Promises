package promises

import "testing"

func TestImmediatelyRejected(t *testing.T) {
	initialReason := "hello"
	promise := Reject(initialReason)
	promise.Then(func(value PromiseResult) Promise {
		t.Error("promise was resolved")
		return Empty()
	}).Fail(func(reason PromiseResult) Promise {
		if reason.(string) != initialReason {
			t.Error("%v != %v", initialReason, reason)
		}
		return Empty()
	})
}
