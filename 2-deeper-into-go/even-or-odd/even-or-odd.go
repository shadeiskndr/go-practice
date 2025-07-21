package main

import (
	"fmt"
)

type myNumbers []int

func newNumbers() myNumbers{
	numbers := myNumbers{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return numbers
}

func (n myNumbers) evenOrOdd(number int) {
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
}

func (n myNumbers) print() {
	for i, number := range n {
		fmt.Println(i, number)
	}
}