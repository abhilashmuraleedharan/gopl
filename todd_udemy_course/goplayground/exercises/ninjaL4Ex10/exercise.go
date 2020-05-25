package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "ice cream", "Sunsets"},
	}
	PrintMap(m)
	fmt.Println("-------------------------------------------------------")

	m["Q"] = []string{"Weapons", "Gadgets", "Super cars"}
	PrintMap(m)
	fmt.Println("-------------------------------------------------------")

	if _, ok := m["Q"]; ok {
		delete(m, "Q")
	}
	PrintMap(m)
}

// PrintMap to print map
func PrintMap(m map[string][]string) {
	for k, v := range m {
		fmt.Println(k, "likes: ")
		for index, item := range v {
			fmt.Printf("\t%d %s\n", index, item)
		}
	}
}
