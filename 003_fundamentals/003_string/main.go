package main

import "fmt"

func main() {
	s := "Hello, World!"
	fmt.Println(s)

	// A string is made up of a slice of byte, so we can do that conversion.
	fmt.Println([]byte(s))
	bs := []byte{72, 101, 108, 108, 111}
	s = string(bs)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(s)

	// We can see the UTF-8 codepoint representation of a character.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println()

	s = `This is a
        multiline
            string`

	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
