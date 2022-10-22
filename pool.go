package main

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

var pool = SafePool{
	pool: make([]*PoolElement, 0),
}

type PoolElement struct {
	Id          string
	Running     bool
	InstancedAt time.Time
	Task
}

type SafePool struct {
	mu   sync.Mutex
	pool []*PoolElement
}

func (safePool *SafePool) push(task Task) PoolElement {
	safePool.mu.Lock()
	defer safePool.mu.Unlock()

	element := PoolElement{
		Id:          uuid.NewString(),
		InstancedAt: time.Now(),
		Running:     false,
		Task:        task,
	}

	safePool.pool = append(safePool.pool, &element)

	return element
}

func (safePool *SafePool) next() *PoolElement {
	safePool.mu.Lock()
	defer safePool.mu.Unlock()

	var targetElement *PoolElement

	for _, element := range safePool.pool {
		if element.Running {
			continue
		}

		if targetElement == nil {
			targetElement = element

			continue
		}

		if targetElement.Task.Proprity > element.Task.Proprity {
			targetElement = element

			continue
		}

		if targetElement.Task.Proprity == element.Task.Proprity && targetElement.InstancedAt.Sub(element.InstancedAt).Milliseconds() < 0 {
			targetElement = element

			continue
		}
	}

	return targetElement
}

func (safePool *SafePool) remove(id string) {
	safePool.mu.Lock()

	for index, element := range safePool.pool {
		if element.Id != id {
			continue
		}

		safePool.pool = append(safePool.pool[:index], safePool.pool[index+1:]...)

		break
	}

	safePool.mu.Unlock()
}
