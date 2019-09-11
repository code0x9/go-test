package main

import (
	"fmt"
	"log"
	"time"
)

const TickDuration = 2 * time.Second

type Adder struct {
	sum   int
	timer *time.Timer
}

func NewAdder() *Adder {
	return &Adder{
		sum:   0,
		timer: time.NewTimer(TickDuration),
	}
}

func (a *Adder) Enqueue(in int) {
	a.sum += in
	a.timer.Reset(TickDuration)
}

func (a *Adder) Flush() int {
	<-a.timer.C
	defer func() {
		a.sum = 0
		//a.timer = time.NewTimer(TickDuration)
	}()
	return a.sum
}

func main() {
	adder := NewAdder()
	log.Printf("please input numbers. all numbers are queued to adder & print added results when idle for %v\n", TickDuration)

	go func() {
		for {
			var in int
			if _, err := fmt.Scanf("%d", &in); err != nil {
				panic(err)
			}
			adder.Enqueue(in)
		}
	}()

	for {
		if sum := adder.Flush(); sum > -1 {
			log.Println("sum", sum)
		}
	}
}
