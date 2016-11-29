package main

//import "fmt"

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func(){
		ch1 <- 1
		ch2 <- 2
	}()

	<-ch2
	<-ch1
}
