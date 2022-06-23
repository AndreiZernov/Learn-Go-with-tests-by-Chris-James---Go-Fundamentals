package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("expected a runs 5 times and equal aaaaa", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("expected a runs 3 times and equal aaa", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		assertCorrectMessage(t, repeated, expected)
	})
}
