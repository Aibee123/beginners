package main

import (
		"fmt"
)

func main() {
		func() {
				msg := "Hello Go!"
				fmt.Println(msg)
		}()
}