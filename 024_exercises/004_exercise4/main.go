package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (e sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", e.lat, e.long, e.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0,
			sqrtError{
				lat:  "123.456",
				long: "567.890",
				err:  fmt.Errorf("invalid input: %v", f),
			}
	}

	return math.Sqrt(f), nil
}
