package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func citire(data string) float64 {
	num, err := strconv.ParseFloat(data, 64)
	if err != nil {
		fmt.Println("Eroare la conversie")
		return -404.0
	}
	return num
}

func main() {
	roboti := make([]Robot, 0)
	razaTest := 7.0

	input, err := os.Open("input.in")
	if err != nil {
		fmt.Println("Eroare la citire")
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()

	n := int(citire(scanner.Text()))

	for i := 0; i < n; i++ {
		scanner.Scan()
		positions := strings.Fields(scanner.Text())

		x := citire(positions[0])
		y := citire(positions[1])

		roboti = append(roboti, Robot{
			Position: Point{X: x, Y: y},
			Activ:    true,
		})
	}
	lcmAveraging(roboti, razaTest)
}
