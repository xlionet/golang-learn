[TOP]

#golang 的 死锁

在用goroutine 的时候经常会碰到死锁的情况，这里做一下总结

## 死锁的四个必要条件
-  `互斥条件` 一个资源每次只能被一个进程使用。
- `请求与保持条件` 一个进程因请求资源而阻塞时，对已获得的资源保持不放。
- `不剥夺条件` 进程已获得的资源，在末使用完之前，不能强行剥夺。
- `循环等待条件` 若干进程之间形成一种头尾相接的循环等待资源关系。
## channel 原理
向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。从
channel中读取数据的语法是
value := <-ch
如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel
中被写入数据为止。我们之后还会提到如何控制channel只接受写或者只允许读取，即单向
channel

##实例
1.只在单一的goroutine里操作无缓冲信道，一定死锁。比如你只在main函数里操作信道:
```go
func main() {
        ch := make(chan int)
        ch <- 1
        fmt.Printf(ch)
}
```
2.循环等待
ch2 等待 ch1 取操作， 而ch1 等待 ch2 写操作
```go
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

```
3.对于channel来说，只流进不流出 或 只流出不流进 都会造成死锁
```go
c, quit := make(chan int), make(chan int)

go func() {
   c <- 1  // c通道的数据没有被其他goroutine读取走，堵塞当前goroutine
   quit <- 0 // quit始终没有办法写入数据
}()

<- quit // quit 等待数据的写
```
4. 但是也不是绝对的
```go
func main() {
        var ch = make(chan int)
    go func()
    {
        ch <- 1 
    }()
}

go 启动了线程， 但是主线程结束了
```

## 避免
1.正确的写数据 和 取数据
```go
ch1, ch2 := make(chan int), make(chan int)

go func() {
    ch1 <- 1
    ch2 <- 0
}()

<- ch1 // 取走c的数据！
<- ch2
```
2.设置缓存信道
```go
ch = make(chan int, 1)
这样，放入一个数据ch 不会挂起当前线程， 再放一个才会挂起
```

## 数据到来的顺序 FIFO
1.非缓冲的情况下，开出多个goroutine
```go
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

result:
0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 
````
2.缓冲信道的情况下，使用单个doroutine
```go
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
1,2
```

##多goroutine的使用总结
1.但缓冲阻塞主线程
````go
var ch chan int = make(chan int)

func foo(i int){
        ch <- i
}
func main() {
     count := 100
     for i:=0; i<count; i++{
        go foo(i)
     }
     
     for i:=0; i<100; i++{
        fmt.Println(<-ch)
     }
}
````

2 . 使用缓冲
```go
var ch chan int = make(chan int,100)
func foo(){
        for i:=0; i<100; i++
        {
            ch<-i
        }
        close(ch)
}
func main() {
        go foo()
        for v := range(ch)
        {
                fmt.Printl(v)
        }
}
````