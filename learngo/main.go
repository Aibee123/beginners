package main

import (
		"fmt"
)

func main() {
		d := divide(5.0, 3.0)
		fmt.Println(d)
}

func divide(a, b float64) float64 {
		return a / b
}