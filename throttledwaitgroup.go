//package throttledwaitgroup
package main

import (
	"sync"
)

type TWG struct {
	Size    int
	current chan struct{}
	wg      sync.WaitGroup
}

func New(throttle int) *TWG {
	size := 4 //chosen by fair dice roll guaranteed to be random.
	if throttle > 0 {
		size = throttle
	}

	return &TWG{
		Size:    size,
		current: make(chan struct{}, size),
		wg:      sync.WaitGroup{},
	}

}

func (twg *TWG) Add() {
	twg.current <- struct{}{}
	twg.wg.Add(1)
}

func (twg *TWG) Done() {
	<-twg.current
	twg.wg.Done()
}

func (twg *TWG) Wait() {
	twg.wg.Wait()
}
