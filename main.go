package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 500; i++ {
		go showMenssage(fmt.Sprintf("This is message number %d", i))
	}

	time.Sleep(2 * time.Second) 
}

func showMenssage(message string) {
	fmt.Println(message)
}
