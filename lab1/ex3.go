package main

import "fmt"

func main() {
	list := []int{4, 13, 78, 16, 2, 78, 99, 105, 19083, 202, 76}

	evenNumbers := []int{}

	for index := range list {
		if list[index]%2 == 0 {
			evenNumbers = append(evenNumbers, list[index])
		}
	}

	fmt.Println(evenNumbers)
}
