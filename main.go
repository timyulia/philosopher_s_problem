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

func host(permission, completion chan bool) {
	num := 0
	for {
		if num == 0 {
			permission <- true
			num++
		}
		if num == 1 {
			select {
			case permission <- true:
				num++
			case <-completion:
				num--
			}
		}

		if num == 2 {
			<-completion
			num--
		}

	}
}

func eat(ph *philo, ind int, wg *sync.WaitGroup, completion, permission chan bool) {
	for i := 0; i < 3; i++ {
		<-permission
		ph.rightChop.Lock()
		ph.leftChop.Lock()
		fmt.Printf("starting to eat %d \n", ind+1)
		time.Sleep(2 * time.Second)
		fmt.Printf("finishing eating %d \n", ind+1)
		ph.rightChop.Unlock()
		ph.leftChop.Unlock()
		completion <- true
	}
	wg.Done()

}

func main() {
	completion := make(chan bool)
	permission := make(chan bool)
	table := make([]*philo, 5)
	sticks := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(chopstick)
	}
	for i := 0; i < 5; i++ {
		table[i] = &philo{sticks[i], sticks[(i+1)%5]}
	}

	wg := sync.WaitGroup{}
	wg.Add(5)

	for j := 0; j < 5; j++ {
		go host(permission, completion)
		go eat(table[j], j, &wg, completion, permission)
	}

	wg.Wait()

}
