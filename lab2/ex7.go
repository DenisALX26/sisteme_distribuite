package main

import (
	"fmt"
	"sync"
)

type Pozitie struct {
	x int
	y int
}

func traseuRobot1(canal chan<- Pozitie, wg *sync.WaitGroup, pozitiiCelalatRobot chan Pozitie) {
	defer wg.Done()

	canal <- Pozitie{1, 1}
	canal <- Pozitie{2, 2}
	close(canal)

	for pozitie := range pozitiiCelalatRobot {
		fmt.Println("Robot 1 are: ", pozitie)
	}
}

func traseuRobot2(canal chan<- Pozitie, wg *sync.WaitGroup, pozitiiCelalatRobot chan Pozitie) {
	defer wg.Done()

	canal <- Pozitie{10, 10}
	canal <- Pozitie{11, 11}
	close(canal)

	for pozitie := range pozitiiCelalatRobot {
		fmt.Println("Robot 2 are: ", pozitie)
	}
}

func main() {
	wg := sync.WaitGroup{}

	canal1 := make(chan Pozitie, 2)
	canal2 := make(chan Pozitie, 2)

	wg.Add(2)

	go traseuRobot1(canal1, &wg, canal2)
	go traseuRobot2(canal2, &wg, canal1)

	wg.Wait()
	fmt.Println("S-au terminat robotii")
}
