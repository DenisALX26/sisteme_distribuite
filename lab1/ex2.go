package main

import "fmt"

func main() {
	list := []int{5, 2, 9, 1, 5, 6}

	for i := range list {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}

	fmt.Println("Lista sortata:", list)
}
