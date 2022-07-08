package main

import (
	"fmt"

	"github.com/mbeaver502/LearningGolang/027_testing_benchmarking/003_examples/acdc"
)

func main() {
	fmt.Println(acdc.Sum([]int{1, 2, 3, 4, 5}...))
	fmt.Println(acdc.Sum(6, 7, 8, 9, 0, 1, 2, 3, 4, 5))
}
