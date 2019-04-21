package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getsum(x int, y int) int {
	return x + y
}

func main() {
	var name = "weiqi"
	fmt.Println(greeting(name))
	fmt.Println(getsum(1, 5))
}
