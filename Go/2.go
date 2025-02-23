// Task 2)

/*
take user input as string
convert string to int
write fibonacci sequence till that
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var in string
	fmt.Printf("Please enter the destination number:")
	fmt.Scanf("%s", &in)
	intege, _ := strconv.Atoi(in)
	var a, b = 0, 1               // a & b are backbone of fibonacci sequence
	for i := 0; i < intege; i++ { // i = No. of steps
		fmt.Println(a)
		a, b = b, a+b
	}
}
