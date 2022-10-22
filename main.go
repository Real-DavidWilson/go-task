package main

import (
	"fmt"
	"time"
)

func main() {
	AddTask(&Task{
		Proprity: 2,
		TaskHandler: func() {
			fmt.Println("TASK 01 PRIORITY 2")
		},
	})

	AddTask(&Task{
		Proprity: 1,
		TaskHandler: func() {
			fmt.Println("TASK 02 PRIORITY 1")
		},
	})

	InitProcessor(ProcessorOptions{
		maxConcurrency: 5,
		interval:       time.Millisecond * 5,
	})
}
