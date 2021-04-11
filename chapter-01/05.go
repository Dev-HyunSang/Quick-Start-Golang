package main

import "fmt"

func main() {
	a := 2
	//b := 5
	b := &a
	// &는 메모리의 주소를 볼 수 있습니다.
	fmt.Println(&a, &b)
	fmt.Println(a, &b)
	fmt.Println(*b)

}
