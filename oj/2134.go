package main 

import "fmt"

type selem byte

const size = 50

type Stack struct {
	data [size]selem 
	top int 
}

func initStack() *Stack{
	stack := new(Stack)
	stack.top = -1
	return stack 
}

func (this * Stack) isEmpty() bool {
	if  this.top == -1 {
		return true 
	} 
	return false 
}

func (this * Stack) push(elem selem) bool{

	if this.top > size - 1 {
		return false 
	}

	this.top++ 
	this.data[this.top] = elem

	return true 
}

func (this * Stack) pop() bool{

	if this.top == -1 {
		return false 
	}

	this.top -- 

	return true
}

func main() {

	var str string 

	stack := initStack()

	flag := false 


	for {

		_, err := fmt.Scanf("%s", &str)

		if err != nil{
			return
		}

		for _, ch := range str {

			if  ch == '(' || ch == '[' || ch == '{' {
				stack.push(selem(ch))
				flag = true 
			}

			if !stack.isEmpty() {

				if (ch == ')' && stack.data[stack.top] == '(') || ( ch == '}' && stack.data[stack.top] == '{') || (ch == ']' && stack.data[stack.top] == '[') {
					stack.pop()
				}

			}
		}


		if flag && stack.isEmpty() {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}

	}
}
