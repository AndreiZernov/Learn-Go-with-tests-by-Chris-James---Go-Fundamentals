package iteration

import "testing"
import "fmt"

func ExampleRepeat() {
	str := Repeat("a", 5)
	fmt.Println(str)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	checkRepeter := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("expected a runs 5 times and equal aaaaa", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		checkRepeter(t, repeated, expected)
	})

	t.Run("expected a runs 3 times and equal aaa", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		checkRepeter(t, repeated, expected)
	})
}
