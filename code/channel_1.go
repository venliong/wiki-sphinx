package main 

import (
	"fmt"
)

var numChan chan int 

func sheep(i int, num chan int) {
	for ; ; i += 2 {
		num <- i 
	}
}

func main() {
	numChan = make(chan int)
	go sheep(1, numChan)
	go sheep(2, numChan)

	for {
		fmt.Println( <- numChan )
	}
}
