package main

import "fmt"

type Integer int

func (a Interge)Less(b Interge) bool  {
	return a < b
}

func (a *Interge)add(b Interge) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	add(b Integer)

}
