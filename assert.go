package assert

import (
	"reflect"
	"testing"
)

// F represents a testing function.
type F func(*testing.T)

// C creates a test case for the given name and testing function.
func C(name string, tf F) F {
	return func(t *testing.T) {
		t.Helper()
		t.Run(name, func(t *testing.T) { t.Helper(); tf(t) })
	}
}

// Apply the given *testing.T object to a testing function as a helper.
func Apply(t *testing.T, tf F) { t.Helper(); tf(t) }

// All combines the given testing functions into a single testing function.
func All(tfs ...F) F {
	return func(t *testing.T) {
		t.Helper()
		for _, tf := range tfs {
			Apply(t, tf)
		}
	}
}

// True expects the given condition to be true.
func True(cond bool) F {
	return func(t *testing.T) {
		t.Helper()
		if !cond {
			t.Fatal("expected true")
		}
	}
}

// False expects the given condition to be false.
func False(cond bool) F {
	return func(t *testing.T) {
		t.Helper()
		if cond {
			t.Fatal("expected false")
		}
	}
}

// NoError expects the given error to be nil.
func NoError(err error) F {
	return func(t *testing.T) {
		t.Helper()
		if err != nil {
			t.Fatalf("\nunexpected error: %s", err.Error())
		}
	}
}

// IsError expects the given error to be set.
func IsError(err error) F {
	return func(t *testing.T) {
		t.Helper()
		if err == nil {
			t.Fatalf("expected error")
		}
	}
}

// Equal expects the given values to be equal.
func Equal(v, e interface{}) F {
	return func(t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(v, e) {
			t.Fatalf("\nexpected: %#v\n  actual: %#v", e, v)
		}
	}
}
