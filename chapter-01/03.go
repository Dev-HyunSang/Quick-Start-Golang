package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpperTwo(name string) (lenght int, uppercase string)  {
	defer fmt.Println("I'm done!")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 3))
	// _는 무시되는 변수임. => 무시된 상태로 컴파일
	totalLenght, upperName := lenAndUpper("HyunSang")

	fmt.Println(totalLenght, upperName)
	repeatMe("HyunSang", "Rin", "Nico")
	fmt.Println(lenAndUpperTwo("HyunSang"))
}