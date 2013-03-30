package main 

import "fmt"	

func main() {
	var num int 
	fmt.Scanf("%d", &num)

	for i := 1; i <= num  ; i++ {
		fmt.Printf("%d*%d=%d\n", num, i, num*i)
	}
}