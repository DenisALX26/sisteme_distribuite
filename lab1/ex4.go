package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Introdu un index ca sa elimini elementul de pe acea pozitie:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	index, _ := strconv.Atoi(input)
	list := []float32{3.0, 15.6, 98.0, 67.2, 6.8, 2.1}

	newList := make([]float32, 0, len(list)-1)
	newList = append(newList, list[:index]...)
	newList = append(newList, list[index+1:]...)

	fmt.Println(newList)
}
