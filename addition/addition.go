package main

import "fmt"

func addition(x int, y int) int  {
	 var result = x + y
	 return result
}

func main() {
	 var result = addition (6, 4)
	 fmt.Println(result)
}