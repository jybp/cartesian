package cartesian_test

import (
	"reflect"
	"testing"

	"github.com/jybp/cartesian"
)

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
