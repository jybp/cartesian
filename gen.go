package cartesian

type Iterator func() (v interface{}, hasNext bool)

// Bool implememts Iterator for Bool
func Bool() Iterator {
	first := true
	return func() (interface{}, bool) {
		if first {
			first = true
			return true, true
		}
		return false, false
	}
}
