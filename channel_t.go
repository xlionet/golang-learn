/**
* @brief: A demo for channel
* @author: zhangxiaojie
* @date: Nov 18, 2016
 */
package main

import "fmt"

func Count(ch chan int)  {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([] chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int, 10)
		//println(i)
		go  Count(chs[i])
	}

 	for _ , ch := range(chs){
		<-ch
		//println(i)
	}
}

/**
* @result: copy from console
C:/Go\bin\go.exe run D:/workspace/go-learn/ZXJ_Demos/channel_t.go
Counting
Counting
Counting

Process finished with exit code 0

* @analise: The given result whit printing three "Count" but Ten, which is beyond expection.
 I thk the reason is :
 	go Count(chs[i])
 	{
 		ch <- 1
 		...
 	}
 	after ch <- 1 finished, the channel of `ch` is not blocked anymore, which means <-ch got stream, by doing this,
 	main thread finished before some of sub-thread made by "go Count'
 */