package main

import "fmt"

func main() {

	slice := []string{"Hello", "world", "!"}

	for i, element := range slice { //   Range will return index and element
		fmt.Println(i, element, ":")

		for _, ch := range element { //   Will print out each letter (%q prints rune, not just a bit)
			fmt.Printf(" %q\n", ch)
		}
	}

}
