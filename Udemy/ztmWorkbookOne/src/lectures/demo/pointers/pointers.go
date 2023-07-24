package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {

	counter := Counter{}

	greet := "Hello"
	world := "World"

	fmt.Println(greet, world)

	replace(&greet, "Hi", &counter)
	fmt.Println(greet, world)

	phrase := []string{greet, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

}
