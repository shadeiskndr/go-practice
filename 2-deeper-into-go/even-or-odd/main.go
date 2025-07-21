package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := newNumbers()
	// for _, number := range numbers {
	// 	numbers.evenOrOdd(number)
	// }

	//Scanner for user input
	scanner := bufio.NewScanner(os.Stdin)
 	fmt.Print("Enter a number: ")
 	scanner.Scan()
 	input := scanner.Text()
 	number, err := strconv.Atoi(input)
 	if err != nil {
 		fmt.Println("Please enter a valid number")
 		return
 	}
	numbers.evenOrOdd(number)
} 
