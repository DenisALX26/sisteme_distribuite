package main

import (
	"math/rand"
	"sync"
)

type LamportClock struct {
	counter int
}

func (lc *LamportClock) Increment() {
	lc.counter++
}

func (lc *LamportClock) Send() {
	lc.Increment()
}

func (lc *LamportClock) Receive(receivedCounter int) {
	if receivedCounter > lc.counter {
		lc.counter = receivedCounter
	}
	lc.Increment()
}

type Message struct {
	sender, ts int
}

type Event struct {
	process, internalCounter int
	kind                     string
}

func less(a, b Event) bool {
	if a.internalCounter != b.internalCounter {
		return a.internalCounter < b.internalCounter
	}
	return a.process < b.process
}

func process(id int, rounds int, next chan<- Message, prev <-chan Message, events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := LamportClock{counter: 0}

	for i := 0; i < rounds; i++ {
		interval := rand.Intn(3)

		for j := 0; j < interval; j++ {
			clock.Increment()
			events <- Event{process: id, internalCounter: clock.counter, kind: "internal"}
		}

		clock.Send()
		next <- Message{sender: id, ts: clock.counter}
		events <- Event{process: id, internalCounter: clock.counter, kind: "send"}

		msg := <-prev
		clock.Receive(msg.ts)
		events <- Event{process: id, internalCounter: clock.counter, kind: "receive"}
	}
}

func main() {
	var wg sync.WaitGroup
	events := make(chan Event)

	ch12 := make(chan Message, 1)
	ch23 := make(chan Message, 1)
	ch31 := make(chan Message, 1)

	wg.Add(3)
	go process(1, 5, ch12, ch31, events, &wg)
	go process(2, 5, ch23, ch12, events, &wg)
	go process(3, 5, ch31, ch23, events, &wg)

	go func() {
		wg.Wait()
		close(events)
	}()

	var eventList []Event
	for event := range events {
		eventList = append(eventList, event)
	}

	for i := 0; i < len(eventList); i++ {
		for j := i + 1; j < len(eventList); j++ {
			if !less(eventList[i], eventList[j]) {
				eventList[i], eventList[j] = eventList[j], eventList[i]
			}
		}
	}

	for _, event := range eventList {
		println("Process:", event.process, "Counter:", event.internalCounter, "Kind:", event.kind)
	}
}
