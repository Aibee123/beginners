package main

import(
		"fmt"
)

func main() {
		number := 50
		guess := 105
		if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
		} else  if guess > 100 {
				fmt.Println("The guess must be less than 100!")
		} else	{
				if guess < number {
						fmt.Println("Too low")
				}
				if guess > number {
						fmt.Println("Too high" )
				}
				if guess == number {
						fmt.Println("You got it")
				}
				fmt.Println(number<=guess, number >=guess, number!=guess)
		}		
}