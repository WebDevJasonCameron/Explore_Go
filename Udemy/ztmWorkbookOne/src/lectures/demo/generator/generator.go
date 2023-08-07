package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()
	return out
}

func generateRandIntn(count, min, max int) <-chan int {

	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randInt := generateRandInt(1, 100)

	fmt.Println("GenerateRandInt infinite")

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("GenerateRandInt using finite count")

	randIntnRange := generateRandIntn(2, 1, 10)

	for i := range randIntnRange {
		fmt.Println(i)
	}

	fmt.Println("GenerateRandInt using finite count and determine if open")

	randIntn := generateRandIntn(3, 1, 10)

	for {
		n, open := <-randIntn
		if !open {
			if !open {
				break
			}
			fmt.Println(n)
		}
	}

}
