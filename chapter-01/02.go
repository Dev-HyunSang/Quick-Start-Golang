package main

import "fmt"

func main() {
	const name = "HyunSang"
	// name = "HyunSang Park"
	fmt.Println(name)

	// 축약형은 오로지 func 안에서만 변수에만 적용 가능
	// names := "HyunSang"
	var names string = "HyunSang"
	names = "Rin"
	fmt.Println(names)

}
