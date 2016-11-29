package main


/**
*	set the `timeout ` = 10s in redis-server.windows.conf file
*	test :if a client is idle for 10s, the redis-server will close it
*	Avail() will get the count
*/



import (
	"github.com/mediocregopher/radix.v2/pool"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func CreatePool1() (p *pool.Pool) {
	p, err := pool.New("tcp", "localhost:6379", 3)
	if err != nil {
		panic(err)
	}

	//wg.Done()
	return
}

func DoCmd1(p *pool.Pool) {

	if p == nil {
		fmt.Println("pool is nil")
		return
	}

	conn1, err := p.Get()
	if err != nil {
		panic(err)
	}

	if conn1.Cmd("SET", "111", 111).Err != nil {
		fmt.Println("err.")
	}

	fmt.Println(p.Avail())

	conn2, err := p.Get()
	if err != nil {
		panic(err)
	}

	if conn2.Cmd("SET", "222", 222).Err != nil {
		fmt.Println("err.")
	}

	fmt.Println(p.Avail())

	conn3, err := p.Get()
	if err != nil {
		panic(err)
	}

	if conn3.Cmd("SET", "333", 333).Err != nil {
		fmt.Println("err.")
	}

	fmt.Println(p.Avail())

	time.Sleep(time.Second * 15) //after  15s three conns will be disable

	p.Put(conn1)
	p.Put(conn2)
	p.Put(conn3)

	fmt.Println(p.Avail())

	conn, err := p.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Avail())

	if conn.Cmd("SET", "aaa",   15).Err != nil {
		fmt.Println("err.")
		p.Put(conn)
		fmt.Println(p.Avail())

	}
	wg.Done()
}

func main() {

p :=CreatePool1()

wg.Add(1)
go DoCmd1(p)

wg.Wait()

	}

/**
*
C:/Go\bin\go.exe run D:/workspace/go-learn/ZXJ_Demos/radix_pool_t.go
15 s passed
p.Avail() is :  2
err.
30 s passed
p.Avail() is :  2
err.

分析 设置了timeout=10s
10s之内，从pool中申请的线程没有执行动作，redis-server便将这个线程回收，因为15s后执行cmd操作出现了错误，
但是p.Avail()依然是2，说明这个线程并没有被回收，那么这个坏的线程去哪了？ 答案，需要执行put（）操作把这个连接
归还给pool

那如果归还的是坏的，下次get的时候能用吗？

见radix_timeout_t.go

*/