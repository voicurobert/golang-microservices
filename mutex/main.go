package main

import (
	"fmt"
	"sync"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *AtomicInt) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

func (i *AtomicInt) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

var (
	counter       = 0
	lock          sync.Mutex
	atomicCounter = AtomicInt{}
)

func updateCounter(wg *sync.WaitGroup) {
	//lock.Lock()
	//defer lock.Unlock()
	//counter++
	atomicCounter.Increase()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	//fmt.Printf("final counter %d\n", counter)
	fmt.Printf("final atomic counter %d\n", atomicCounter.Value())
}
