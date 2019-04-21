package main

import "fmt"

func replace(a *int) { //replace by point that point to address
	fmt.Println(a)
	*a = 10
}

func main() {

	a := 6
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	replace(&a)
	fmt.Println(a)
}
