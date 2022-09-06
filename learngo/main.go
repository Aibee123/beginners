package main

import(
		"fmt"
)

func main () {
		var students [3] string
		fmt.Printf("Students: %v\n", students)
		students[0] = "Lisa"
		students[1] = "Arnold"
		students[2] = "Ahmed"
		fmt.Printf("Students: %v\n", students)
}