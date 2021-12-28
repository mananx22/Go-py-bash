// Task 1) make a variable and echo it out while simultaneously commenting out this line of comment.

/*
make single line comment
make a block comment
make a variable
print the variable on screen
make a global variable
overrride global variable from local variable
*/

package main

import (
	"fmt"
)

var global = "manan tunsa" //we can remove data type
var global2 int = 25

func main() {
	var global2 int = 52
	fmt.Println("Welcome global2 from main function is -> ", global2)
	fmt.Println("Welcome global from function is -> ", global)
}
