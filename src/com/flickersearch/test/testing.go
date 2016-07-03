package test

import (
	"reflect"
	"runtime/debug"
	"testing"
)

// AssertEquals asserts if the two objects are the same.
func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		debug.PrintStack()
		t.Errorf("Expected: %s, got: %s\n", expected, actual)
	}
}

// AssertNil asserts if the given object is nil.
func AssertNil(t *testing.T, actual interface{}) {
	val := reflect.ValueOf(actual)
	if !val.IsNil() {
		debug.PrintStack()
		t.Errorf("Expected: nil, got: %s\n", actual)
	}
}

// AssertNotNil asserts if the given object is not nil.
func AssertNotNil(t *testing.T, actual interface{}) {
	val := reflect.ValueOf(actual)
	if val.IsNil() {
		debug.PrintStack()
		t.Errorf("Expected: not nil, got: %s\n", actual)
	}
}


// AssertKeyFoundInMap checks if a key is a valid key.
func AssertKeyFoundInMap(t *testing.T, mapType interface{}, key interface{}) {
	m := reflect.ValueOf(mapType)
	if !m.MapIndex(reflect.ValueOf(key)).IsValid() {
		debug.PrintStack()
		t.Errorf("Expected key %s in the map", key)
	}
}


// AssertError asserts if error is set.
func AssertError(t *testing.T, err error) {
	if err == nil {
		debug.PrintStack()
		t.Fatal("Expected error")
	}
}

// AssertNotError asserts if error is not set.
func AssertNotError(t *testing.T, err error) {
	if err != nil {
		debug.PrintStack()
		t.Fatalf("Expected no error, got %s instead", err)
	}
}

// Fail fails the test.
func Fail(t *testing.T, message ...interface{}) {
	debug.PrintStack()
	t.Fatal(message)
}
