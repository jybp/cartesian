package cartesian_test

import (
	"reflect"
	"testing"

	"github.com/jybp/cartesian"
)

func assertIterator(t *testing.T, expected []interface{}, iterator cartesian.Iterator) {
	actual := []interface{}{}
	for {
		v, hasNext := iterator()
		actual = append(actual, v)
		if !hasNext {
			break
		}
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}

func TestBool(t *testing.T) {
	expected := []interface{}{
		true, false,
	}
	assertIterator(t, expected, cartesian.Bool())
}

func TestOne(t *testing.T) {
	expected := []interface{}{
		"a",
	}
	assertIterator(t, expected, cartesian.One("a"))
}

func TestAll(t *testing.T) {
	expected := []interface{}{
		"a", 1,
	}
	assertIterator(t, expected, cartesian.From("a", 1))
}
