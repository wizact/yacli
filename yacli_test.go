package yacli

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	a := NewApplication()

	if got, want := a.Description, "cli tool"; got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
