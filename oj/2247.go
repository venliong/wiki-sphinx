package main 

import "fmt"

func main() {
	var str string 
	var num rune

	fmt.Scanf("%s %c", &str, &num)

	var j int = 0
	for _, i := range str {
		if num == i {
			j ++ 
		}
	}

	fmt.Println(j)
}