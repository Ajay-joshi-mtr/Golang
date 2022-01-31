package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

var counter = 0
var m = sync.RWMutex{} //avoid race condition and sync

func main() {
	runtime.GOMAXPROCS(100) //set threads cores
	fmt.Printf("Threads : %v\n", runtime.GOMAXPROCS(-1))

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()

	}
	wg.Wait()
	/*
		destroyed concurrency and force single threded way using mutex
	*/

}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// os threads having their stack or thread pooling  --expensive in other lang
// in go different like ER-lang called Green thread

/* advices
1. Don't create goroutine in libraries
2. When creating Go routines, Know how to end (Avoid memory leaks)
3. Avoid race condition
*/
