package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	x, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Eroare la conversia inputului:", err)
		return
	}

	if x < 2 {
		fmt.Println("Nu este numar prim")
		return
	}

	if x == 2 {
		fmt.Println("Este numar prim")
		return
	}

	if x%2 == 0 {
		fmt.Println("Nu este numar prim")
		return
	}

	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			fmt.Println("Nu este numar prim")
			return
		}
	}

	fmt.Println("Este numar prim")
	return
}
