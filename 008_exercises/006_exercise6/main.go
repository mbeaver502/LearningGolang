package main

import "fmt"

func main() {
	states := make([]string, 50)

	x := []string{
		`Alaska`,
		`Alabama`,
		`Arkansas`,
		`Arizona`,
		`California`,
		`Colorado`,
		`Connecticut`,
		`Delaware`,
		`Florida`,
		`Georgia`,
		`Hawaii`,
		`Iowa`,
		`Idaho`,
		`Illinois`,
		`Indiana`,
		`Kansas`,
		`Kentucky`,
		`Louisiana`,
		`Massachusetts`,
		`Maryland`,
		`Maine`,
		`Michigan`,
		`Minnesota`,
		`Missouri`,
		`Mississippi`,
		`Montana`,
		`North Carolina`,
		`North Dakota`,
		`Nebraska`,
		`New Hampshire`,
		`New Jersey`,
		`New Mexico`,
		`Nevada`,
		`New York`,
		`Ohio`,
		`Oklahoma`,
		`Oregon`,
		`Pennsylvania`,
		`Rhode Island`,
		`South Carolina`,
		`South Dakota`,
		`Tennessee`,
		`Texas`,
		`Utah`,
		`Virginia`,
		`Vermont`,
		`Washington`,
		`Wisconsin`,
		`West Virginia`,
		`Wyoming`}

	fmt.Println(len(states))
	fmt.Println(cap(states))

	//for i, s := range x {
	//	states[i] = s
	//}

	copy(states, x)

	for i, s := range states {
		fmt.Println(i, s)
	}
}
