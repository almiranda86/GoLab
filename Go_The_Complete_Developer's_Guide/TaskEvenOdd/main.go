package main

import "fmt"

func main() {
	number := newEvenOdd()

	for _, n := range number {
		if n%2 == 0 {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
	}
}

func newEvenOdd() []int {
	var numbers []int

	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range values {
		numbers = append(numbers, v)
	}

	return numbers
}
