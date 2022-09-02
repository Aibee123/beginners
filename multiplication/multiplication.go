package main

import "fmt"

func multiplication(x int , y int) int {
	var result = x * y
	return result
}

func main()  {
	var result = multiplication(8, 6)
	fmt.Println(result)
}