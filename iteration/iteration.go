package iteration

import "fmt"

func ExampleRepeat() {
	str := Repeat("a", 5)
	fmt.Println(str)
	// Output: aaaaa
}

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
