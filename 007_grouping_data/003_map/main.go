package main

import "fmt"

func main() {
	// A map is a key-value store, like a dict.
	// map[keyType]valueType
	m := map[string]int{
		"James":      32,
		"Moneypenny": 24,
	}

	fmt.Println(m)

	// We can add key-value pairs to a map.
	m["John"] = 20
	fmt.Println(m)

	// We can iterate over the key-value pairs in a map.
	for key, value := range m {
		fmt.Println(key, value)
	}

	// We can delete key-value pairs from a map.
	// If the key to be deleted does not exist, then the delete is effectively a nop with no error.
	delete(m, "John")
	fmt.Println(m)

	// If a key isn't in the map, then you'll get back the zero value for the value type.
	fmt.Println(m["Steve"])

	// We can check that a key exists in the map by using the comma-ok idiom.
	v, ok := m["Steve"]
	fmt.Println(v, ok)

	// We can use the comma-ok idiom as part of a conditional.
	if _, ok := m["Steve"]; !ok {
		fmt.Println("Steve does not exist in the map.")
	}
}
