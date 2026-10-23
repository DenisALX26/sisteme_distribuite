package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	beginExercise(1)
	scanner.Scan()
	input := scanner.Text()
	x, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Eroare la conversia inputului:", err)
		return
	}

	if !ex1(x) {
		fmt.Println("Nu este numar prim")
	} else {
		fmt.Println("Este numar prim")
	}

	beginExercise(2)
	listaEx2 := []int{5, 2, 9, 1, 5, 6}
	sortedList := ex2(listaEx2)
	fmt.Println("Lista sortata:", sortedList)

	beginExercise(3)
	listaEx3 := []int{4, 13, 78, 16, 2, 78, 99, 105, 19083, 202, 76}
	evenNumbers := ex3(listaEx3)
	fmt.Println("Lista cu numere pare:", evenNumbers)

	beginExercise(4)
	fmt.Print("Introdu un index ca sa elimini elementul de pe acea pozitie:")
	scanner.Scan()
	input = scanner.Text()
	index, _ := strconv.Atoi(input)
	listEx4 := []float32{3.0, 15.6, 98.0, 67.2, 6.8, 2.1}
	newList := ex4(listEx4, index)
	fmt.Println("Lista cu elementul eliminat este: ", newList)

	beginExercise(5)
	listEx5 := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	fmt.Print("Lista initiala: ", listEx5)
	newStringList := ex5(listEx5)
	fmt.Println("Lista finala: ", newStringList)

	beginExercise(6)
	ex6()

	beginExercise(7)
	ex7()

	beginExercise(8)
	ex8()

	beginExercise(9)
}
