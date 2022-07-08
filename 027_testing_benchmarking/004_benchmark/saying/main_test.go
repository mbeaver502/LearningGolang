package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")

	if s != "Welcome, James" {
		t.Error("got", s, "expected", "Welcome, James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))

	// Output: Welcome, James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
