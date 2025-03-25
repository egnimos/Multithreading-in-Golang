package main

import (
	"sync"
	"time"
)

var (
	amount = 100
	lock   = sync.Mutex{}
)

func withdraw() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		amount -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("completed withdrawing")
}

func credit() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		amount += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("completed crediting")
}

func main() {
	for i := 0; i < 10; i++ {
		go credit()
		go withdraw()
		// println(amount)
		time.Sleep(5000 * time.Millisecond)
		println(amount)
	}
}
