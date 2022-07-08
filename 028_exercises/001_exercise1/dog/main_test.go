package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func ExampleYears() {
	age := Years(10)
	fmt.Println(age)

	// Output: 70
}

func TestYears(t *testing.T) {
	type test struct {
		data int
		want int
	}

	tests := []test{
		{
			data: 1,
			want: 7,
		},
		{
			data: 2,
			want: 14,
		},
		{
			data: 7,
			want: 49,
		},
		{
			data: 10,
			want: 70,
		},
	}

	for _, test := range tests {
		if got := Years(test.data); got != test.want {
			t.Error("want", test.want, "got", got)
		}
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYearsTwo() {
	age := YearsTwo(10)
	fmt.Println(age)

	// Output: 70
}

func TestYearsTwo(t *testing.T) {
	type test struct {
		data int
		want int
	}

	tests := []test{
		{
			data: 1,
			want: 7,
		},
		{
			data: 2,
			want: 14,
		},
		{
			data: 7,
			want: 49,
		},
		{
			data: 10,
			want: 70,
		},
	}

	for _, test := range tests {
		if got := YearsTwo(test.data); got != test.want {
			t.Error("want", test.want, "got", got)
		}
	}
}
