
package main

import "fmt"
//The should be a demo for chan`s capacity
func WriteStr(ch chan string) {
	for  {
		ch <- "abcddsadasdasdasdasdasdasdasdasdasda"
	}

}

func WriteNum(ch chan string) {
	for{
		ch <- "1234"
	}
}

func Read(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}

func WriteLine(ch chan string){
	ch <- "write a line"
}

func main()  {
	ch := make(chan string)
	//go WriteNum(ch)
	//go WriteStr(ch)
	 WriteLine(ch)
	println("hello")
	//Read(ch)
}

