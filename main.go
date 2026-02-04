package main

import (
	"sync"
	"time"
)

func main() {
	var fabio sync.Mutex
	i := 0
	for j := 0; j < 10000; j++ {
		go func() {
			fabio.Lock()
			i++
			fabio.Unlock()
		}()
	}
	time.Sleep(time.Second * 2)
	println(i)
}

func changeNumber(i *int, number int) {
	*i = number
}
