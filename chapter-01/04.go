package main

import "fmt"

// for, range, args
func superAdd(numbers ...int) int {
	for number := range numbers {
		fmt.Println(number)
	}
	return 1
}

func superAddTwo(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func superAddThree(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	const annotation string = "========================="
	superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(annotation)
	superAddTwo(1, 2, 3, 4, 5, 6)
}
