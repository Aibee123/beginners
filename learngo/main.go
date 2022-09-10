package main

import (
		"fmt"
)

func main() {
		var i interface {} = [3]int{}
		switch i. (type) {
		case int:
				fmt.Println("i is an int")
		case float64:
				fmt.Println("i is a float64")
		case string:
				fmt.Println("i is a string")
		case [3]int:
				fmt.Println("i is [3]int")
			default:
				fmt.Println("i is another type")
		}
}