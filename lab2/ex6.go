package main

import (
	"fmt"
	"time"
)

func robotCursa(nume string, canal chan<- string) {
	if nume == "Robot 1" {
		time.Sleep(2 * time.Second)
		canal <- "Robot 1 a terminat cursa"
	} else {
		time.Sleep(1 * time.Second)
		canal <- "Robot 2 a terminat cursa"
	}
}

func main() {
	finishRobo1 := make(chan string)
	finishRobo2 := make(chan string)

	go robotCursa("Robot 1", finishRobo1)
	go robotCursa("Robot 2", finishRobo2)

	select {
	case status := <-finishRobo1:
		fmt.Println(status)
	case status := <-finishRobo2:
		fmt.Println(status)
	}
}
