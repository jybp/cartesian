package cartesian_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jybp/cartesian"
)

func ExampleProduct() {
	sets := cartesian.Product([]interface{}{"a", "b", "c"}, []interface{}{1, 2, 3})
	for _, set := range sets {
		fmt.Println(set)
	}

	// Ordered Output:
	// [a 1]
	// [b 1]
	// [c 1]
	// [a 2]
	// [b 2]
	// [c 2]
	// [a 3]
	// [b 3]
	// [c 3]
}

func TestProduct(t *testing.T) {
	input := [][]interface{}{
		{"a1"},
		{"b1", "b2"},
		{"c1", "c2", "c3"},
		{"d1"},
	}

	expected := [][]interface{}{
		{"a1", "b1", "c1", "d1"},
		{"a1", "b1", "c2", "d1"},
		{"a1", "b1", "c3", "d1"},
		{"a1", "b2", "c1", "d1"},
		{"a1", "b2", "c2", "d1"},
		{"a1", "b2", "c3", "d1"},
	}

	actual := cartesian.Product(input...)

	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}
