package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m[`Joker`] = []string{`Hahaha`, `Why so serious`, `All part of the plan`}
	for i, v := range m {
		fmt.Println("Record for : ", i)
		for j, v1 := range v {
			fmt.Println(j, v1)
		}
		fmt.Println()
	}

}
