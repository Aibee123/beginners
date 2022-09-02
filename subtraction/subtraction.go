package main

import "fmt"

func subtraction(x int, y int) int {
	var result = x - y
	return result
}

func main() {
	var result = subtraction(5, 3)
	fmt.Println(result)
}