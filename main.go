package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Producer is type
type Producer struct {
	OutChan chan int
}

func (p *Producer) getOutChan() <-chan int {
	return p.OutChan
}

func (p *Producer) produce() {
	for {
		time.Sleep(3 * time.Second)
		p.OutChan <- rand.Int()
	}
}

func main() {
	p := Producer{
		OutChan: make(chan int, 10),
	}

	go p.produce()

	prodChan := p.getOutChan()
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case a := <-prodChan:
			fmt.Println("Go message from producer", a)
		case <-ticker.C:
			fmt.Println("Go message from tiker")
		}
	}
}
