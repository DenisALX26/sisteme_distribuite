package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c *Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimeaza(time1, time2, time3, time4 time.Time) (theta, delta time.Duration) {
	theta = (time2.Sub(time1) + time3.Sub(time4)) / 2
	delta = (time4.Sub(time1) - time3.Sub(time2)) / 2
	return theta, delta
}

func masoara(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	time1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	time2 := s_clock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	time3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	time4 := c_clock.Now()

	return estimeaza(time1, time2, time3, time4)
}

func main() {
	offset := 2 * time.Second
	s_clock := &Clock{offset: 0}

	// punctul a
	c_clock := &Clock{offset: offset}
	theta_a, delta_a := masoara(c_clock, s_clock)
	c_clock.offset += delta_a
	error_a := c_clock.offset.Abs()

	fmt.Println("theta = ", theta_a, "delta = ", delta_a)
	fmt.Println("Transformare offset: ", offset, " -> ", c_clock.offset)

	if error_a < offset.Abs() {
		fmt.Println("A scazut offset-ul\n\n")
	} else {
		fmt.Println("Nu a scazut offset-ul\n\n")
	}

	// punctul b
	c_clock = &Clock{offset: offset}
	min_theta, min_delta := masoara(c_clock, s_clock)

	for i := 0; i < 19; i++ {
		theta, delta := masoara(c_clock, s_clock)
		fmt.Printf("%d: theta = %v, delta = %v\n", i+1, theta, delta)

		if delta < min_delta {
			min_theta = theta
			min_delta = delta
		}
	}

	c_clock.offset += min_delta
	error_b := c_clock.offset.Abs()

	fmt.Println("theta = ", min_theta, "delta = ", min_delta)
	fmt.Println("Transformare offset: ", offset, " -> ", c_clock.offset)

	if error_b < offset.Abs() {
		fmt.Println("A scazut offset-ul\n\n")
	} else {
		fmt.Println("Nu a scazut offset-ul\n\n")
	}
}
