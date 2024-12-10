package signal

import "testing"

func TestHandler(t *testing.T) {
	t.Log("here is a plain message")
	t.FailNow()
	t.Logf("here is my message %d", 123)
	// t.Error("failing test")
}
