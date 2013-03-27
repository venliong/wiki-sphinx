package main

import "fmt"

func main() {
	var str string 
	var ch byte 

	fmt.Scanf("%s", &str)
	fmt.Scanf("%c", &ch)

	oldStrByte := []byte(str)
	var newStrByte []byte

	for _, c := range oldStrByte {
		if  ch != c {
			newStrByte = append(newStrByte, c)
		}
	}

	fmt.Println(string(newStrByte))
}