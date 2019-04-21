package main

import ("fmt"
		"math"
		"github.com/weiqi/go_crash_course/03_package/strutil"
	)

func main(){
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Pow(2,6))
	fmt.Println(strutil.Reverse("olleh"))
}