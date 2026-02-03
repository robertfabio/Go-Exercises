package main

import (
	"fmt"
	"os"
)

func ReadFile(){
	_, err := os.Open("non_existent_file.txt")
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("File opened successfully")
	}
}

func main() {
	
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	ReadFile()

	fmt.Println("OI???????")
}