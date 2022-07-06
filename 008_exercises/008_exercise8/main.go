package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_doctor":       {"Evil", "Ice Cream", "Sunsets"},
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)

		for i, x := range v {
			fmt.Println(i, x)
		}
	}
}
