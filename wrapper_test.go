package main

import (
	"fmt"
	"log"
)

func ExampleNormalTask() {
	err := taskWrapper(normalTask)
	if err != nil {
		log.Println(err)
	}
	// Output:
	// before
	// ... Running Task ...
	// after
}

func ExampleFailingTask() {
	err := taskWrapper(failingTask)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// before
	// ... Task Failure ...
	// after
	// TASK ERROR
}

func taskWrapper(task func() error) error {
	defer after()
	if err := before(); err != nil {
		return err
	}

	return task()
}

func normalTask() error {
	fmt.Println("... Running Task ...")
	return nil
}

func failingTask() error {
	fmt.Println("... Task Failure ...")
	return fmt.Errorf("TASK ERROR")
}

func before() error {
	fmt.Println("before")
	return nil
}

func after() {
	fmt.Println("after")
}
