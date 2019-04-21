package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loops though ids
	for i, id := range ids { //index, value
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	sum := 0
	for _, id := range ids { //if not use index, put _ on first parameter
		sum += id
	}

	fmt.Println(sum)

	//range for MAP
	emails := map[string]string{"wang": "wang@gmail.com", "weiqi": "weiqi@gmail.com"}
	for k, v := range emails { //if not use index, put _ on first parameter
		fmt.Printf("%s : %s\n", k, v)
	}
}
