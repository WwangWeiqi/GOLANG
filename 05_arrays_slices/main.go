package main

import "fmt"

func main() {
	//Array
	// var fruitArr [2]string
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	//declare and assign
	fruitArr := []string{"apple", "orange"}
	fruitArr = append(fruitArr, "banana")
	fmt.Println(fruitArr, len(fruitArr))
	fmt.Println(fruitArr[1:2])

	emptyArr := []string{}
	fmt.Println(len(emptyArr))
}
