package main

import "fmt"

var ch1 = make(chan string)
var ch2 = make(chan string)

func do(s string) {
	ch1 <- <-ch2
	fmt.Println(ch1)
}

func main() {
	go do("hello")
	<-ch1
}
