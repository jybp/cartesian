package cartesian

type Iterator func() (v interface{}, hasNext bool)

// One returns an Iterator with one value
func One(v interface{}) Iterator {
	return func() (interface{}, bool) { return v, false }
}

// All returns an Iterator that iterates throught "all"
// "all" length should be at least one.
func All(all ...interface{}) Iterator {
	if len(all) == 0 {
		return func() (interface{}, bool) { return nil, false }
	}
	idx := 0
	return func() (interface{}, bool) {
		v := all[idx]
		last := idx == len(all)-1
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
