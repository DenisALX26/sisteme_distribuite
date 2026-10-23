package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	s, _ := strconv.Atoi(input[1])

	scanner.Scan()
	input = strings.Split(scanner.Text(), " ")

	list := make([]int, len(input))
	pos := make(map[int]int)
	for i := range list {
		list[i], _ = strconv.Atoi(input[i])
		pos[list[i]] = i
	}

	for i := range list {
		target := s - list[i]
		_, ok := pos[target]
		if ok && i < pos[target] {
			fmt.Println(i, pos[target])
		}
	}
}
