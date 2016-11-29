package main

import "fmt"

var ch chan int = make(chan int, 2)

func main() {
	ch <- 1
	ch <- 2
	close(ch) // without this ,there will be a deadlock error
	for v := range ch{
		fmt.Println( v)
	}

}
