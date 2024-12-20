package main

import "fmt"

func evenOrOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		fmt.Println(number, evenOrOdd(number))
	}
}
