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
	p2 := new(int) // <=> var p2 *int
	p2 = &b
	fmt.Println(p2)
	fmt.Printf("%T", p2)

	//zero value
	var pointer1 *int
	pointer2 := new(int)

	fmt.Println(pointer1)
	fmt.Println(pointer2)

	a1 := 100
	var pointer3 *int

	pointer3 = &a1
	*pointer3 = 999 // a := 999

	fmt.Println(pointer3)
	fmt.Println(a1)
}
