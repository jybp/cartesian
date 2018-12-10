package cartesian

// Product implementation by Paul Hankin https://stackoverflow.com/a/29004530
func Product(sets ...[]interface{}) [][]interface{} {
	lens := func(i int) int { return len(sets[i]) }
	product := [][]interface{}{}
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []interface{}
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

// ForAll evaluates fn for all possible iterators values.
// "fn" has to be a function with the same number of parameters as the provided iterators.
// "fn" may return a simple bool to stop the evaluations (false means to stop).
func ForAll(fn interface{}, iterators ...Iterator) error {
	return nil
}
