package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["weiqi"] = "442056022@gmail.com"
	emails["wang"] = "442056022@sina.com"
	emails["Mike"] = "MIKE@sina.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	delete(emails, "wang")

	fmt.Println(emails)
	fmt.Println(len(emails))

	emptyMap := map[string]string{}
	fmt.Println(len(emptyMap))
}
