package main 

import "fmt"	

func main() {
	var N int 
	fmt.Scanf("%d", &N)

	sum := 0 
	max := 0
	min := 1000

	for i := 0; i< N; i++ {
		var j int 
		fmt.Scanf("%d", &j)
		sum = sum + j

		if j > max {
			max = j
		}

		if j < min  {
			min = j
		}
	}

	fmt.Printf("%d %d %d\n", max, min, sum/N)
}