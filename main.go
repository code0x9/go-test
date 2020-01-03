package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//adder()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
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
