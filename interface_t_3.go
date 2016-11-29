package main

import (
	"fmt"
)

type PeopleGetter interface {
	GetName() string
	GetAge() int
}

type EmployeeGetter interface {
	PeopleGetter
	GetSalary() int
	Help()
}

type Employee struct {
	name string
	age int
	salary int
	gender string
}

func (self *Employee) GetName() string {
	return self.name
}

func (self *Employee) GetAge() int {
	return self.age
}

func (self *Employee) GetSalary() int {
	return self.salary
}

func (self *Employee) Help() {
	fmt.Println("This is help infomations.")
}

type Man struct {
	gender interface{
		GetGender() string
	       }
}

func (self *Employee) GetGender() string {
	return self.gender
}

type Callback interface {
	Execute()
}

type CallbackFunc func()

func (self CallbackFunc) Execute() {self()}

func main() {
	var varEmptyInterface interface{}
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)
	varEmptyInterface = 100
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)
	varEmptyInterface = "GoLang"
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	varEmployee := Employee{
		name:	"Passi",
		age:	24,
		salary:	999,
		gender:	"Male",
	}
	fmt.Println("varEmployee is: ", varEmployee)
	varEmployee.Help()
	fmt.Println("varEmployee.name", varEmployee.GetName())
	fmt.Println("varEmplyee.age", varEmployee.GetAge())
	fmt.Println("varEmployee.salary ", varEmployee.GetSalary())

	varMan := Man{&Employee{
		name:	"Nobody",
		age:	20,
		salary:	10000,
		gender:	"Unkown",
	}}
	fmt.Println("The gender of Nobody is: ",varMan.gender.GetGender())

	var varEmpInter EmployeeGetter = &varEmployee
	switch varEmpInter.(type) {
	case nil:
		fmt.Fprintln("nil")
	case PeopleGetter:
		fmt.Println("PeopleGetter")
	default:
		fmt.Println("Unkown")
	}

	varCallbacker := CallbackFunc(func() {println("I am a Callback function")})
	varCallbacker.Execute()
}




















