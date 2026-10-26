package main

import (
	"fmt"
)

func primulCanal(c chan string) {
	c <- "Mesaj canal 1"
}

func alDoileaCanal(c chan string) {
	c <- "Mesaj canal 2"
}

func main() {
	msg := make(chan string)

	go primulCanal(msg)
	go alDoileaCanal(msg)

	primulMesaj := <-msg
	alDoileaMesaj := <-msg

	fmt.Printf("Primit de pe canal: %s\n", primulMesaj)
	fmt.Printf("Primit de pe canal: %s\n", alDoileaMesaj)
}
