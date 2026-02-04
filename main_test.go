package main

import "testing"

func TestGetenv(t *testing.T) {
	got := getenv("DOES_NOT_EXIST", "x")
	if got != "x" {
		t.Fatalf("expected default, got %q", got)
	}
}
