package cartesian

type Iterator func() (v interface{}, hasNext bool)

// From returns an Iterator that iterates throught "values"
// "values" length should be at least one.
func From(values ...interface{}) Iterator {
	if len(values) == 0 {
		return func() (interface{}, bool) { return nil, false }
	}
	idx := 0
	return func() (interface{}, bool) {
		v := values[idx]
		last := idx == len(values)-1
		idx++
		return v, !last
	}
}

// Bool returns an Iterator for bool
func Bool() Iterator {
	first := true
	return func() (interface{}, bool) {
		if first {
			first = false
			return true, true
		}
		return false, false
	}
}
