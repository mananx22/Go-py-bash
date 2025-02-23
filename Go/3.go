// Task 3)

/*
make a swap function.
pass by value first.
prove that original variables are still unaffected.
pass by pointer/reference again.
prove that original variables are indeed affected.
*/

package main

import "fmt"

func main() {
	var x, y int
	x = 4
	y = 8
	fmt.Printf("Initial value of x is: %d & value of y is: %d \n", x, y)

	swapbyvalue(x, y)
	fmt.Printf("value of x is: %d & value of y is: %d after passing values only. \n", x, y)

	swapbypointer(&x, &y) // using & operator represents "the address of"
	fmt.Printf("value of x is: %d & value of y is: %d after passing pointers. \n", x, y)

}

func swapbyvalue(a, b int) {
	a, b = b, a
}

func swapbypointer(a, b *int) {
	*a, *b = *b, *a
}
