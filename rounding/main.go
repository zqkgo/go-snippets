package main

import "fmt"

func main() {
	// rounding to n
	n := 32
	fmt.Println(31 &^ (n - 1)) // 0
	fmt.Println(32 &^ (n - 1)) // 32
	fmt.Println(55 &^ (n - 1)) // 32
	fmt.Println(70 &^ (n - 1)) // 64
}
