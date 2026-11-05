package main

import (
	"fmt"
	"sync"
)

type Node struct {
	id     int
	left   chan int
	right  chan int
	status string
}

func (n *Node) execute(rounds int, wg *sync.WaitGroup) {
	defer wg.Done()

	send_id := n.id

	for i := 0; i < rounds; i++ {
		n.right <- send_id
		recv_id := <-n.left

		if recv_id == n.id {
			n.status = "LEADER"
			fmt.Println("Leader ID: ", n.id)
			send_id = -1
		} else if recv_id > n.id {
			send_id = recv_id
		} else {
			send_id = -1
		}
	}
}

func main() {
	var ws sync.WaitGroup

	nodes := []Node{}

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	nodes = append(nodes, Node{id: 1, left: ch3, right: ch1, status: "non-lider"})
	nodes = append(nodes, Node{id: 2, left: ch1, right: ch2, status: "non-lider"})
	nodes = append(nodes, Node{id: 3, left: ch2, right: ch3, status: "non-lider"})

	rounds := 3

	for i := range nodes {
		ws.Add(1)
		go nodes[i].execute(rounds, &ws)
	}

	ws.Wait()
}
