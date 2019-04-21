package main

import "fmt"

func replace(a *int) int {
	*a = 10
	return *a
}

func main() {

	a := 6
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	replace(&a)
	fmt.Println(a)
}
