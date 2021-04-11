package main

import "fmt"

func main() {
	names := []string{"hyunsang", "rin", "dal"}
	names = append(names, "flyrin")
	fmt.Println(names)
}
