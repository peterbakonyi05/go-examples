// Requirements:
// Implement the dining philosopher’s problem with the following constraints/modifications.
// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// The host allows no more than 2 philosophers to eat concurrently.
// Each philosopher is numbered, 1 through 5.
// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.\

// implementation notes:
// - mutex is used for shared resource (chopstick)
// - waitgroup is used to know when everything finished
// - channel is used to limit the concurrency

package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const maxConcurrentEaters = 2
const numPhilosophersEating = 3

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p Philosopher) eat(hostPermission chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numPhilosophersEating; i++ {
		<-hostPermission // Wait for host permission

		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		time.Sleep(time.Second) // Simulate eating time
		fmt.Printf("finishing eating %d\n", p.number)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		hostPermission <- struct{}{} // Release host permission
	}
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
		}
	}

	var wg sync.WaitGroup
	hostPermission := make(chan struct{}, maxConcurrentEaters)

	for i := 0; i < maxConcurrentEaters; i++ {
		hostPermission <- struct{}{}
	}

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].eat(hostPermission, &wg)
	}

	wg.Wait()
}
