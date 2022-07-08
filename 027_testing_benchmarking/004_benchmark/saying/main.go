package saying

import "fmt"

func Greet(s string) string {
	return fmt.Sprintf("Welcome, %v", s)
}
