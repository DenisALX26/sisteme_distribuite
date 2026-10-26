package main

import "fmt"

func mesaj2(msg chan string) {
	mesajPrimit := <-msg
	fmt.Println("Mesajul primit de la main: " + mesajPrimit)
}

func main() {
	msg := make(chan string)

	go mesaj2(msg)
	msg <- "salut"

	fmt.Println("S-a terminat main()")
}
