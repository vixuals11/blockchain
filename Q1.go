package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func PersonD(x Person) {
	fmt.Println("Name: ", x.name)
}

func main() {
	var P1 Person
	P1.name = "HH"
	P1.age = 90
	P1.job = "Student"
	P1.salary = 200

	PersonD(pers1)

}
