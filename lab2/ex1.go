package main

import (
	"fmt"
)

func mesaj(msg chan string) {
	msg <- "salut"
	fmt.Println("Se trmite mesajul catre main")
}

func main() {
	test := make(chan string)

	go mesaj(test)

	msg := <-test
	fmt.Printf("Mesaj primit din alt goroutine: \"%s\"\n", msg)
}
