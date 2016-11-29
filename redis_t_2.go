package main

import(
//"github.com/mediocregopher/radix.v2/redis"
)
import (
	"github.com/mediocregopher/radix.v2/redis"
	"time"
	"fmt"
)

func main(){
	client, err := redis.DialTimeout("tcp", "localhost:6379", time.Duration(0))
	if err != nil {
		panic(err)
	}
	//err = client.Cmd("set", "bar","qweqw", "xx").Err
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//r := client.Cmd("MGET", "foo", "bar", "baz")
	//if r.Err != nil {
	//	panic(r.Err)
	//}


	/* Pipelining */

	client.PipeAppend("GET", "foo")
	client.PipeAppend("SET", "bar", "foo")
	client.PipeAppend("DEL", "baz")

	foo, err := client.PipeResp().Str()
	if err != nil {

	}
	fmt.Println(foo)
	if err := client.PipeResp().Err; err != nil {

	}

	if err := client.PipeResp().Err; err != nil {

	}


	/**

	Pipelining is when the client sends a bunch of commands to the server at once,
	and only once all the commands have been sent does it start reading the replies off the socket.
	 This is supported using the PipeAppend and PipeResp methods.
	 PipeAppend will simply append the command to a buffer without sending it,
	 the first time PipeResp is called it will send all the commands in the buffer and return the Resp
	 for the first command that was sent. Subsequent calls to PipeResp return Resps for subsequent commands:

	 以上是出自https://godoc.org/github.com/mediocregopher/radix.v2/redis的官方解释，自己试验后给出的解释如下：
	 该驱动是基于tcp协议与主机通信，cmd指令能够及时的将这条指令发给redis执行并返回结果，而PipeAppend则是将所有的指令一次性执行
	 并将每一条的结果保存到一条队列里，然后使用PipeResp()去获得返回结果
	 */


	/* Flatten    扁平化?*/

	client.Cmd("HMSET", "myhash", "key1", "val1", "key2", "val2")
	client.Cmd("HMSET", "myhash", []string{"key1", "val1", "key2", "val2"})
	client.Cmd("HMSET", "myhash", map[string]string{
		"key1": "val1",
		"key2": "val2",
	})
	client.Cmd("HMSET", "myhash", [][]string{
		[]string{"key1", "val1"},
		[]string{"key2", "val2"},
	})




}