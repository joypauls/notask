package main

import (
	"testing"
)

func TestPadding(t *testing.T) {
	// test
	got := padding("test")
	want := " test "
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
