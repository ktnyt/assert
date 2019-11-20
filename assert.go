package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// F represents a testing function.
type F func(testing.TB)

// Nop does nothing.
func Nop(tb testing.TB) {}

// Eval runs the given function and returns a Nop.
func Eval(f func()) F { f(); return Nop }

// C creates a test case for the given name and testing function.
func C(name string, tfs ...F) F {
	return func(tb testing.TB) {
		if v, ok := tb.(interface{ Helper() }); ok {
			v.Helper()
		}
		switch v := tb.(type) {
		case *testing.T:
			v.Run(name, func(t *testing.T) { t.Helper(); Apply(t, tfs...) })
		case *testing.B:
			v.Run(name, func(b *testing.B) { b.Helper(); Apply(b, tfs...) })
		default:
			panic(fmt.Errorf("%T is not *testing.T nor *testing.B", v))
		}
	}
}

// Apply the given testing.TB object to testing functions as a helper.
func Apply(tb testing.TB, tfs ...F) {
	if v, ok := tb.(interface{ Helper() }); ok {
		v.Helper()
	}
	for _, tf := range tfs {
		tf(tb)
	}
}

// All combines the given testing functions into a single testing function.
func All(tfs ...F) F {
	return func(tb testing.TB) {
		if v, ok := tb.(interface{ Helper() }); ok {
			v.Helper()
		}
		for _, tf := range tfs {
			tf(tb)
		}
	}
}

// True expects the given condition to be true.
func True(cond bool) F {
	return func(tb testing.TB) {
		if v, ok := tb.(interface{ Helper() }); ok {
			v.Helper()
		}
		if !cond {
			tb.Fatal("expected true")
		}
	}
}

// False expects the given condition to be false.
func False(cond bool) F {
	return func(tb testing.TB) {
		if v, ok := tb.(interface{ Helper() }); ok {
			v.Helper()
		}
		if cond {
			tb.Fatal("expected false")
		}
	}
}

// NoError expects the given error to be nil.
func NoError(err error) F {
	return func(tb testing.TB) {
		if v, ok := tb.(interface{ Helper() }); ok {
			v.Helper()
		}
		if err != nil {
			tb.Fatalf("\nunexpected error: %s", err.Error())
		}
	}
}

// IsError expects the given error to be set.
func IsError(err error) F {
	return func(tb testing.TB) {
		if v, ok := tb.(interface{ Helper() }); ok {
			v.Helper()
		}
		if err == nil {
			tb.Fatalf("expected error")
		}
	}
}

// Equal expects the given values to be equal.
func Equal(v, e interface{}) F {
	return func(tb testing.TB) {
		if v, ok := tb.(interface{ Helper() }); ok {
			v.Helper()
		}
		if !reflect.DeepEqual(v, e) {
			tb.Fatalf("\nexpected: %#v\n  actual: %#v", e, v)
		}
	}
}
