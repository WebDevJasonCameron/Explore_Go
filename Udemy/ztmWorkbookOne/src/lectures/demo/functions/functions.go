package main

import "fmt"

func myFunc(word string) string {
	return word
}

func main() {

	word := myFunc("a word")

	fmt.Println("I took in: ", word)

}
