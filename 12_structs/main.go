package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

//greeting method (value receiver)
func (p Person) greeting() string {
	return "Hello my name is " + p.firstName + p.lastName + " I am " + strconv.Itoa(p.age)
}

//grow age (pointer receiver)
func (p *Person) hasBirthday() {
	fmt.Println("p in function", p)
	p.age++
}

func main() {
	Person1 := Person{"weiqi", "wang", "Palo Alto", "M", 27}

	fmt.Println(Person1.firstName)

	fmt.Println(Person1.greeting())
	Person1.hasBirthday()
	fmt.Println(Person1.greeting())
}
