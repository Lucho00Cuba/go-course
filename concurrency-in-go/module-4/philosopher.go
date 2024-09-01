package main

import (
	"fmt"
	"sync"
	"time"
)

type Stick struct {
	sync.Mutex
}

type Ps struct {
	num, count          int
	leftChop, rightChop *Stick
}

func (p Ps) eat(c chan *Ps, wait *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- &p
		if p.count < 3 {
			p.leftChop.Lock()
			p.rightChop.Lock()

			fmt.Println("starting to eat ", p.num)
			p.count = p.count + 1
			fmt.Println("finishing eating", p.num)

			p.rightChop.Unlock()
			p.leftChop.Unlock()
			wait.Done()
		}

	}
}

func host(c chan *Ps, wait *sync.WaitGroup) {
	for {
		if len(c) == 2 {
			<-c
			<-c
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func start() []*Ps {
	ChopSticks := make([]*Stick, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(Stick)
	}

	Philosophers := make([]*Ps, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Ps{i + 1, 0, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	return Philosophers
}

func main() {
	var wait sync.WaitGroup
	c := make(chan *Ps, 2)

	wait.Add(15)

	P := start()

	go host(c, &wait)
	for i := 0; i < 5; i++ {
		go P[i].eat(c, &wait)
	}
	wait.Wait()
}
