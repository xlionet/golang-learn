package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person) printMsg() {
	fmt.Printf("I am %s, and my age is %d\n", p.name, p.name)
}

func (p Person) eat(s string) {
	fmt.Printf("I am %s, and i am eating %s", p.name, s)
}

func (p Person) drink(s string) {
	fmt.Printf("I am %s, and i am drinking %s", p.name, s)
}

type People interface {
	printMsg()
	PeopleEat()
	PeopleDrink()
	//eat()
}

type PeopleDrink interface {
	drink(s string)
}

type PeopleEat interface {
	eat(s string)
}

type PeopleEatDrink interface {
	eat()
	drink()
}

type Foodie struct {
	name string
}

func (f Foodie) eat(s string)  {
	fmt.Printf("I am foodie, %s my favorite food is the %s", f.name, s)
}

func main() {
	var p1 People
	p1 = Person{"Rain", 23}
	p1.printMsg()
	p1.drink("orange juice")
}