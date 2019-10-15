package main

import "testing"

func TestServer(t *testing.T) {
	want := "Hello, world."
	if got := Server(); got != want {
		t.Errorf("Server() = %q, want %q", got, want)
	}
}
