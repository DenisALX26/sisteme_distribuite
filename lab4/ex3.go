package main

import (
	"fmt"
	"sync"
)

type NodeEx3 struct {
	id    int
	value float64
	in    []chan float64
	out   []chan float64
}

func (n *NodeEx3) execute(rounds int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < rounds; i++ {
		for _, ch := range n.out {
			ch <- n.value
		}

		suma := n.value

		for _, ch := range n.in {
			suma += <-ch
		}

		n.value = suma / float64(len(n.in)+1)
	}

	var decizie bool = false
	if n.value > 0.5 {
		decizie = true
	}
	fmt.Printf("Node %d: medie %.4f cu decizia %t\n", n.id, n.value, decizie)
}

func main() {
	var wg sync.WaitGroup
	nodes := []NodeEx3{}

	nodes = append(nodes, NodeEx3{id: 1, value: 1, in: nil, out: nil})
	nodes = append(nodes, NodeEx3{id: 2, value: 0, in: nil, out: nil})
	nodes = append(nodes, NodeEx3{id: 3, value: 1, in: nil, out: nil})

	muchii := [][2]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}

	for _, muchie := range muchii {
		ch1 := make(chan float64, 1)
		ch2 := make(chan float64, 1)

		nodes[muchie[0]].out = append(nodes[muchie[0]].out, ch1)
		nodes[muchie[1]].in = append(nodes[muchie[1]].in, ch1)

		nodes[muchie[1]].out = append(nodes[muchie[1]].out, ch2)
		nodes[muchie[0]].in = append(nodes[muchie[0]].in, ch2)
	}

	rounds := 3

	for i := range nodes {
		wg.Add(1)
		go nodes[i].execute(rounds, &wg)
	}

	wg.Wait()
}
