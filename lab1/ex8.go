package main

import "fmt"

func main() {
	employees := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	departments := make(map[int][]string)
	for name, dep := range employees {
		departments[dep] = append(departments[dep], name)
	}

	fmt.Println(departments)
}
