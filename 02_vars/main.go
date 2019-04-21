package main

import "fmt"

func main(){
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	var name = "weiqi"
	var age int32 = 33

	//shortend define variable
	isCool := true

	first,last := "weiqi","wang"

	fmt.Println(name)
	fmt.Println(first,last)
	fmt.Printf("%T\n",name)
	fmt.Printf("%T\n",age)
	fmt.Printf("%T\n",isCool)
	
}