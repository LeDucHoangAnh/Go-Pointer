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

	//pointer -> array

	array := [3]int{1, 2, 3}

	var pointerAr *[3]int

	pointerAr = &array

	fmt.Println(pointerAr)
	fmt.Printf("%T", pointerAr)
	fmt.Println()
	//pointer -> func
	c := 40
	var pointerFunc *int = &c
	appPointer(pointerFunc)
	fmt.Println(c)
}

func appPointer(pointer *int) {
	*pointer = 7778
}
