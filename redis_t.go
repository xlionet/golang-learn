package main

import (
	"github.com/mediocregopher/radix.v2/redis"
	"time"
	"fmt"
)

func main() {

	client, err := redis.DialTimeout("tcp", "localhost:6379", time.Second * 10)
	if err != nil {
		// handle err
		panic(err)
	}
	r, err := client.Cmd("GET", "foo").Str()
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
	a , err := client.Cmd("INCR","foo").Str()
	if err != nil {
		panic(err)
	}

	fmt.Println(a)

}
