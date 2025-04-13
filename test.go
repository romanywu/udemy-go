package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

}

func primary(numbers *[]int, transform func(any) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func secondary(multiplier int) func(int) int {
	return func(i int) int {
		return multiplier * i
	}
}
