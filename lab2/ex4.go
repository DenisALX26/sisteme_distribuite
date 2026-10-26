package main

import (
	"fmt"
	"time"
)

func procesare(ok chan bool) {
	fmt.Println("A inceput procesarea")

	time.Sleep(3000 * time.Millisecond)

	fmt.Println("S-a terminat procesarea")
	ok <- true
}

func main() {
	ok := make(chan bool)
	go procesare(ok)
	<-ok
	fmt.Println("Programul s-a terminat")
}
