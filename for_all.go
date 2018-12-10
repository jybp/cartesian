package cartesian

import (
	"fmt"
	"reflect"
)

// ForAll evaluates fn for all possible iterators values.
// "fn" has to be a function with the same number of parameters as the provided iterators.
// "fn" may return a simple bool to stop the evaluations (false means to stop).
func ForAll(fn interface{}, iterators ...Iterator) error {
	fnValue := reflect.ValueOf(fn)
	if err := checkFnType(fnValue, len(iterators)); err != nil {
		return err
	}

	sets := [][]interface{}{}
	for _, it := range iterators {
		set := []interface{}{}
		for {
			v, hasNext := it()
			set = append(set, v)
			if !hasNext {
				break
			}
		}
		sets = append(sets, set)
	}

	cases := Product(sets...)

	for _, cas := range cases {
		values := []reflect.Value{}
		for _, val := range cas {
			values = append(values, reflect.ValueOf(val))
		}
		out := fnValue.Call(values)
		if len(out) == 1 && !out[0].Bool() {
			return nil
		}
	}

	return nil
}

func checkFnType(fnValue reflect.Value, lenIterators int) error {
	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		return fmt.Errorf("fn has to be a func: %v", fnType.Kind())
	}
	if fnType.NumIn() != lenIterators {
		return fmt.Errorf("number of fn parameters does not match number of iterators: %d != %d", fnType.NumIn(), lenIterators)
	}
	if fnType.NumOut() > 1 {
		return fmt.Errorf("fn has more than one return value: %d", fnType.NumOut())
	}
	if fnType.NumOut() == 1 && fnType.Out(0).Kind() != reflect.Bool {
		return fmt.Errorf("fn return value should be bool: %v", fnType.Out(0).Kind())
	}
	return nil
}
