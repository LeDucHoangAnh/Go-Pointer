package main

import "fmt"

func main() {
	//Khai báo 1 con trỏ
	a := 100
	var pointer *int
	pointer = &a
	fmt.Println(pointer)
	fmt.Printf("%T", pointer)

	fmt.Println()
	b := 123
	p2 := new(int)
	p2 = &b
	fmt.Println(p2)
	fmt.Printf("%T", p2)
}
