package errmsg

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	expected := errors.New("something is wrong @ errmsg_test.go:13;\x1e")

	e := New("something is wrong")
	if diff := cmp.Diff(e.Error(), expected.Error()); diff != "" {
		t.Error("Encode(e); (-got +want)", diff)
	}
}

func TestHere(t *testing.T) {
	expected := "errmsg_test.go:22"

	e := Here()
	if diff := cmp.Diff(string(e), expected); diff != "" {
		t.Error("Encode(e); (-got +want)", diff)
	}
}
