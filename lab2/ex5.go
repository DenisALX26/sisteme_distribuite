package main

func send(canal chan<- string, msg string) {
	canal <- msg
}

func receive(input <-chan string, output chan<- string) {
	msg := <-input
	output <- msg
}

func main() {
	canalSendReceive := make(chan string)
	go send(canalSendReceive, "salut")
	go receive(canalSendReceive, canalSendReceive)

	msg := <-canalSendReceive
	println(msg)
}
