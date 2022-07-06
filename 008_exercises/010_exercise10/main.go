package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_doctor":       {"Evil", "Ice Cream", "Sunsets"},
	}

	m["q_agent"] = []string{"Gadgets", "Guns"}
	delete(m, "bond_james")

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)

		for i, x := range v {
			fmt.Println(i, x)
		}
	}
}
