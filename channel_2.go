package  main

import "fmt"

var c chan int =  make(chan int)

func foo(i int){
	c <- i
}

func main() {
	for i := 0; i < 10; i++ {
		go foo(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", i)

	}
}