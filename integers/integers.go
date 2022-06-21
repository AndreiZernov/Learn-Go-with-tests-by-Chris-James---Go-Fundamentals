package integers

import "fmt"

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
