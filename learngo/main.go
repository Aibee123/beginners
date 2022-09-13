package main

import (
		"fmt"
)

func main() {
		var i interface{} = 1
		switch i. (type) {
		case int:
				fmt.Println("i is an interger")
		case string:
				fmt.Println("i is a string")
		default:
				fmt.Println("I don't know what i is")
		}
}