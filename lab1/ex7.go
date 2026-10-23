package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("file.txt")
	scanner := bufio.NewScanner(file)

	counts := make(map[string]int)

	for scanner.Scan() {
		_, ok := counts[scanner.Text()]

		if !ok {
			counts[scanner.Text()] = 1
		} else {
			counts[scanner.Text()]++
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Eroare la citirea fisierului")
	}

	fmt.Println(counts)
}
