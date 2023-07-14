//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//

package main

import "fmt"

// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello,", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func msg() string {
	return "Hello"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer

// * Write a function that returns any number
func aNum() int {
	return 3
}

// * Write a function that returns any two numbers
func twoNums() (int, int) {
	return 2, 4
}

// * Add three numbers together using any combination of the existing functions.
func addThreeNums(n1, n2, n3 int) int {
	return n1 + n2 + n3
}

func main() {

	greet("smash")
	fmt.Println(msg())

	n1, n2 := twoNums()
	answer := addThreeNums(aNum(), n1, n2)
	fmt.Println(answer)

}

// hello?  Test commit...
