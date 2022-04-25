package main

import (
	"fmt"
	"sync"
	"time"
)

type chopstick struct{ sync.Mutex }

type philo struct {
	leftChop  *chopstick
	rightChop *chopstick
}

func eat(ph *philo, ind int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ph.rightChop.Lock()
		ph.leftChop.Lock()
		fmt.Printf("starting to eat %d \n", ind+1)
		time.Sleep(1 * time.Second)
		fmt.Printf("finishing eating %d \n", ind+1)
		ph.rightChop.Unlock()
		ph.leftChop.Unlock()
	}
	wg.Done()

}

func main() {
	table := make([]*philo, 5)
	sticks := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(chopstick)
	}
	for i := 1; i < 5; i++ {
		table[i] = &philo{sticks[i], sticks[(i+1)%5]}
	}
	table[0] = &philo{sticks[1], sticks[0]}

	wg := sync.WaitGroup{}
	wg.Add(5)

	for j := 0; j < 5; j++ {
		go eat(table[j], j, &wg)
	}

	wg.Wait()

}
