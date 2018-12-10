package cartesian_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jybp/cartesian"
)

func ExampleForAll() {
	err := cartesian.ForAll(func(a bool, b int) {
		// Do something with a and b
		fmt.Println(a, b)
	}, cartesian.Bool(), cartesian.From(1, 2))

	if err != nil {
		// Handle error
	}

	// Ordered Output:
	// [true 1]
	// [true 2]
	// [false 1]
	// [false 2]
}

func TestForAll(t *testing.T) {

	expected := [][]bool{
		{true, true},
		{false, false},
		{true, false},
		{false, true},
	}

	actual := [][]bool{}

	err := cartesian.ForAll(func(a, b bool) bool {
		actual = append(actual, []bool{a, b})
		return true
	}, cartesian.Bool(), cartesian.Bool())

	if nil != err {
		t.Fatal(err)
	}

	if 4 != len(actual) {
		t.Fatal(len(actual), "!= 4")
	}

	for _, e := range expected {
		found := false
		for _, a := range actual {
			if reflect.DeepEqual(e, a) {
				found = true
			}
		}
		if !found {
			t.Fatal(e, "not found")
		}
	}
}

// TestForAllPanic documents the current behavior
func TestForAllPanicTypeMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		}
	}()
	cartesian.ForAll(func(a int) {}, cartesian.Bool())
}
