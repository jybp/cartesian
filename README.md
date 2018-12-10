# cartesian
Package to create cartesian products

## Usage

```go
import (
    "fmt"
    "github.com/jybp/cartesian"
)

func main() {

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
```

## ForAll usage

```go
import (
    "fmt"
    "github.com/jybp/cartesian"
)

func main() {
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
```