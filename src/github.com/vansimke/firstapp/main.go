package main

import (
	"fmt"
	"os"
)

func mydetails(name  string, age int, phonenumber string, height float64) string {
	//	ibinabo, 24, +1238148969926, 5'9

	s := fmt.Sprintf("%s is %d years old.\n", name, age, phonenumber, height)
	print(s)
}


