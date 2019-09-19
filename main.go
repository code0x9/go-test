package main

import (
	"fmt"
	"log"
)

func main() {
	adder()
}

func adder() {
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
