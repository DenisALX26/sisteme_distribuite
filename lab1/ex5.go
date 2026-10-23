package main

import "fmt"

func main() {
	list := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	fmt.Print("Lista initiala: ", list)

	i, j := 0, len(list)-1

	for {
		if i < j {
			list[i], list[j] = list[j], list[i]
			i++
			j--
		} else {
			fmt.Println("Lista finala: ", list)
			return
		}
	}
}
