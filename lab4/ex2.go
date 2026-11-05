package main

import (
	"fmt"
	"sync"
)

type NodeEx2 struct {
	id     int
	in     []chan int
	out    []chan int
	status string
}

func (n *NodeEx2) execute(rounds int, wg *sync.WaitGroup) {
	defer wg.Done()

	maxID := n.id

	for i := 0; i < rounds; i++ {
		for _, ch := range n.out {
			ch <- maxID
		}

		for _, ch := range n.in {
			recvID := <-ch
			if recvID > maxID {
				maxID = recvID
			}
		}
	}

	if maxID == n.id {
		n.status = "LEADER"
		fmt.Println("Lider ID: ", n.id)
	}
}

func main() {
	var wg sync.WaitGroup
	nodes := []NodeEx2{}

	nodes = append(nodes, NodeEx2{id: 1, in: nil, out: nil, status: "non-leader"})
	nodes = append(nodes, NodeEx2{id: 2, in: nil, out: nil, status: "non-leader"})
	nodes = append(nodes, NodeEx2{id: 3, in: nil, out: nil, status: "non-leader"})

	muchii := [][2]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}

	for _, muchie := range muchii {
		ch := make(chan int, 1)
		nodes[muchie[0]].out = append(nodes[muchie[0]].out, ch)
		nodes[muchie[1]].in = append(nodes[muchie[1]].in, ch)
	}

	rounds := 3

	for i := range nodes {
		wg.Add(1)
		go nodes[i].execute(rounds, &wg)
	}

	wg.Wait()
}
