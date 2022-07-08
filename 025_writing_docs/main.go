// We can run a localhost server that generates docs for our modules:
// godoc -http=:8080 (or whatever port)

package main

import (
	"fmt"

	"github.com/mbeaver502/LearningGolang/025_writing_docs/mymath"
)

func main() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(mymath.Sum(xi...))
}
