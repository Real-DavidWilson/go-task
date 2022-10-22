package main

import (
	"time"
)

var options ProcessorOptions

type ProcessorOptions struct {
	maxConcurrency int
	interval       time.Duration
}

func InitProcessor(processorOptions ProcessorOptions) {
	options = processorOptions

	for i := 0; i < options.maxConcurrency; i++ {
		go processor()
	}
}

func processor() {
	for {
		time.Sleep(options.interval)

		element := pool.next()

		if element == nil {
			continue
		}

		element.Task.Run()
		pool.remove(element.Id)
	}
}
