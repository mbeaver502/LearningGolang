package main

import "fmt"

func main() {
	s := []string{"Hello"}
	fmt.Println(s)

	s = append(s, "World")
	fmt.Println(s)

	s = append(s, "this", "is", "a", "test")
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for _, v := range s {
		fmt.Println(v)
	}

	// We can slice elements out of a slice by using a range.
	// [inclusive:exclusive]
	fmt.Println(s[2:6])
	fmt.Println(s[:2])
	fmt.Println(s[2:])
	fmt.Println(s[:])

	s = s[2:]
	fmt.Println(s)

	x := []string{"Hello", "World"}
	fmt.Println(x)

	// The ... allows us to unfurl the elements of the s slice so that we can append them to x.
	x = append(x, s...)
	fmt.Println(x)

	// "Delete" elements from a slice by re-slicing on a range.
	// We can also use append to glue separate pieces back together.
	x = append(x[:1], x[4:]...)
	fmt.Println(x)

	// We can use make to create a slice of a given size ahead-of-time.
	// Every time we resize our slice, we're having to create and destroy underlying arrays.
	// To help mitigate that overhead, we can start with a given size already allocated.
	// We can also set a capacity for the underlying array. We don't have to set the capacity.
	// When we exceed capacity, Go will extend the underlying array and update the slice struct.
	z := make([]string, 10, 100)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	// This will throw an error since we're out of range.
	//z[10] = "hi"

	// This will resize the slice without throwing that error.
	z = append(z, "hi")
	fmt.Println(z)

	jb := []string{"James", "Bond", "chocolate", "martini"}
	mp := []string{"Miss", "Moneypenny", "pistachio", "strawberry"}
	fmt.Println(jb)
	fmt.Println(mp)

	// We can have multidimensional slices.
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
