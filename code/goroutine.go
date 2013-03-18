package main 

import "fmt"	


func Add(x,y int) {
	fmt.Println(x+y)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(1, 1)
	}
}