package main

import (
	"fmt"
	"os"
)

func ShowText() {
	fmt.Println("Finalized")
}

func main() {

	
	file, err := os.Create("example.txt")
	defer file.Close()
	
	defer ShowText()
	
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	file.Write([]byte("Fuckoff"))
}