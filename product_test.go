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
