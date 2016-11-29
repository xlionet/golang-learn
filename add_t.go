/**
* @brief: A demo for `type` and type`s function in two manuals: by pointer and by value.
* @author: zhangxiaojie
* @date: Nov 18, 2016
 */
package main

import (
	"fmt"
)

type Interge int

// by pointer
func (a *Interge)add(b Interge ) {
	*a += b
	//fmt.Println(a)
}

//by value
func (a Interge)add1(b Interge ) {
	a += b
	//fmt.Println(a)
}

func main()  {
	a := Interge(5)
	var b Interge = 5
	a.add(2)
	fmt.Println("a = ", a )

	b.add1(2)
	fmt.Println("b = ", b )
}
/**
* @result:
* a = 7
* b = 5
 */
