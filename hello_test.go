package test_moudles

import "testing"

func TestHello(t *testing.T) {
	want := "hello"
	if got := Hello(); got != want {
		t.Errorf("got:%v,want:%v", got, want)
	}
}
