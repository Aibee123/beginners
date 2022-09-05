package main

import (
	"fmt"
)

func mydetails(name  string, age int, phonenumber string, height float64) string {

	result := fmt.Sprintf("my name is %s i am %d years old my phonenumber %s my height is %f", name, age, phonenumber, height)
	return result
}
func main(){
	var result = mydetails("ibinabo", 24, "+1238148969926", 5.8)
	fmt.Println(result)
}
