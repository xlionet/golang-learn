package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyric string) {
	fmt.Println("La la la...", lyric)
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n",e.name, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyric string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-zzzzz"},"MIT", 0.00}
	paul := Student{Human{"Paul", 22, "2121-323-sadd"}, "ahut",879}
	sam := Student{Human{"Sam", 21, "000-323-sadd"}, "ahut",811}

	var i Men
	i = mike
	fmt.Println("This is Mike, a student:")
	i.SayHi()
	i.Sing("Hello, world")

	fmt.Println("Let`s use a slice of men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _,value := range x {
		value.SayHi()
	}
}