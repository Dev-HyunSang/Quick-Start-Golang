package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Bulgogi", "ramen"}
	hyunsang := person{"hyunsang", 18, favFood}
	fmt.Println(hyunsang)
	fmt.Println(hyunsang.name)
}
