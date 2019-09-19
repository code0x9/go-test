package main

import (
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
	}()
	return a.sum
}
