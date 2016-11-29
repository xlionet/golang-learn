/**
* @breif: This go file calculate the qps (The execute time of command per second)of some commonly redis command
*which from http://doc.redisfans.com/index.html
* author: zhangxiaojie
* date : Nov, 24 2016
*/

package main


import (
	"github.com/mediocregopher/radix.v2/redis"
	"time"
	"fmt"
)

type myclient struct{
	cli *redis.Client
}


func (my *myclient) do(cmd string, args ...interface{}) float64{

	start := time.Now().UnixNano()

	var ret *redis.Resp

	for i := 0; i < 10000; i++ {
		ret = my.cli.Cmd(cmd, args)
		if ret.Err != nil {
			fmt.Println("err", cmd, "i", i)
		}
	}

	end := time.Now().UnixNano()

	// Print the result
	fmt.Println(cmd, "finished!")
	if ret.IsType(redis.Str) {
		fmt.Println(ret.Str())

	} else if ret.IsType(redis.Int) {
		fmt.Println(ret.Int())

	}else if ret.IsType(redis.Array) {
		fmt.Println(ret.Array())

	}else {
		fmt.Println("Unknow respose type...")
	}

	s := (float64) (end - start) / 1000000000
	qps := 10000 / s
	return qps

}

type  MyRedis interface {
	do(cmd string, args ...interface{}) float64
}

func main() {
	var mycli myclient
	var err error
	mycli.cli, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return
	}
	var myredis MyRedis = &mycli

	// key
	fmt.Println(myredis.do("DEL", "111"))
	fmt.Println(myredis.do("DUMP", "222"))
	fmt.Println(myredis.do("EXISTS", "111"))
	fmt.Println(myredis.do("Expire", "222", 10))
	fmt.Println(myredis.do("KEYS", "2*"))
	fmt.Println(myredis.do("MOVE", "baz", 2))
	fmt.Println(myredis.do("PERSIST", "333"))
	fmt.Println(myredis.do("RANDOMKEY" ))
	fmt.Println(myredis.do("RENAME", "bar", "newbar"))
	fmt.Println(myredis.do("RESTORE", "333", "dasda"))
	fmt.Println(myredis.do("SORT", "mylist"))
	fmt.Println(myredis.do("TTL", "foo"))
	fmt.Println(myredis.do("TYPE", "foo"))

/*
	//string

	fmt.Println(myredis.do("APPEND", "mykey", "------"))
	fmt.Println(myredis.do("GETRANGE","mykey", 0, 3))
	fmt.Println(myredis.do("SETBIT", "mybit", 0, 1))
	fmt.Println(myredis.do("BITCOUNT", "mybit"))
	fmt.Println(myredis.do("DECR", "444"))
	fmt.Println(myredis.do("DECRBY", "444", 10))
	fmt.Println(myredis.do("GET", "mykey"))
	fmt.Println(myredis.do("GETBIT", "mykey", 1))
	fmt.Println(myredis.do("GETSET","mykey" ,"value"))
	fmt.Println(myredis.do("GETSET", "mykey", "newvalue"))
	fmt.Println(myredis.do("INCR", "num"))
	fmt.Println(myredis.do("INCRBY", "num", 10))
	fmt.Println(myredis.do("INCRBYFLOAT", "num", 1.0))
	fmt.Println(myredis.do("MGET", "a", "b", "c"))
	fmt.Println(myredis.do("MSET", "111", "100", "qqq", "200"))
	fmt.Println(myredis.do("MSETNX", "1", 1, "2", 2))
	fmt.Println(myredis.do("PSETEX", "psetex", 1000, "hello"))
	fmt.Println(myredis.do("SETRANGE", "greeting", 2, "hello, world"))
	fmt.Println(myredis.do("GETRANGE", "greeting", 2, 8))
	fmt.Println(myredis.do("STRLEN", "greeting"))
*/

/*
	// hash
	fmt.Println(myredis.do("GET", "mykey"))
	fmt.Println(myredis.do("SET", "set", "hello"))
	fmt.Println(myredis.do("HSET", "website", "google", "www.google.com"))
	fmt.Println(myredis.do("HDEL", "website", "baidu"))
	fmt.Println(myredis.do("HEXISTS", "website", "baidu"))
	fmt.Println(myredis.do("HMSET", "website", "google", "www.google.com", "yahoo", "www.yahoo.com", "baidu", "www.baidu.com"))
	fmt.Println(myredis.do("HMGET", "website", "google", "yahoo"))
	fmt.Println(myredis.do("HGETALL", "website"))
	fmt.Println(myredis.do("HMSET", "nums", "1", 1, "2", 2))
	fmt.Println(myredis.do("HINCRBY", "nums", "1", 10))
	fmt.Println(myredis.do("HINCRBYFLOAT", "nums", "2", 10.1))
	fmt.Println(myredis.do("KEYS", "website"))
	fmt.Println(myredis.do("HLEN", "website"))
	fmt.Println(myredis.do("HSETNX", "website", "yahoo", "asdasd"))
	fmt.Println(myredis.do("HVALS", "website"))
	//fmt.Println(myredis.do("HSCAN", 0))
*/
	//list
/*
	fmt.Println(myredis.do("LPUSH", "language", "go"))
	fmt.Println(myredis.do("LPUSHX", "language", "python"))
	fmt.Println(myredis.do("RPUSH", "language", "C++"))
	fmt.Println(myredis.do("RPUSH", "language", "node.js"))
	fmt.Println(myredis.do("RPUSH", "language", "node.js"))
	fmt.Println(myredis.do("RPUSH", "language", "php"))
	fmt.Println(myredis.do("RPUSHX", "language", "Java"))
	fmt.Println(myredis.do("LLEN", "language"))
	fmt.Println(myredis.do("LRANGE", "language", 0, 100))
	fmt.Println(myredis.do("LPOP", "language"))
	fmt.Println(myredis.do("LPOP", "language"))
	fmt.Println(myredis.do("BLPOP", "language", 0))
	fmt.Println(myredis.do("DEL", "job", "command", "request"))
	fmt.Println(myredis.do("LPUSH", "command", "update system..."))
	fmt.Println(myredis.do("LPUSH", "request", "visit page"))
	fmt.Println(myredis.do("BLPOP", "job", "command", "request", 0))
	fmt.Println(myredis.do("RPOPLPUSH","language", "receiver"))
	fmt.Println(myredis.do("BRPOPLPUSH", "language", "receiver", 0))
	fmt.Println(myredis.do("LINDEX", "language", -1))
	fmt.Println(myredis.do("LINSERT", "language", "BEFORE", "node.js", "mylanguage"))
	fmt.Println(myredis.do("LREM","language", "1", "node.js"))
	fmt.Println(myredis.do("LPUSH","language", "node.js"))
	fmt.Println(myredis.do("LTRIM", "language", -1, -1))
*/
	//set
/*
	fmt.Println(myredis.do("SADD", "bbs", "discuz.net", "baidu.com"))
	fmt.Println(myredis.do("SADD", "bbs", "worldpress.net"))
	fmt.Println(myredis.do("SADD", "web", "discuz.net", "sina.com"))
	fmt.Println(myredis.do("SCARD", "bbs"))
	fmt.Println(myredis.do("SDIFF", "bbs", "web"))
	fmt.Println(myredis.do("SDIFFSTORE", "bbs", "web", "diff"))
	fmt.Println(myredis.do("SINTER", "bbs", "web"))
	fmt.Println(myredis.do("SINTERSTORE", "bbs", "web", "inter"))
	fmt.Println(myredis.do("SISMEMBER", "bbs", "discuz.net"))
	fmt.Println(myredis.do("SMEMBERS", "bbs"))
	fmt.Println(myredis.do("SMOVE","web", "bbs", "sina.com"))
	fmt.Println(myredis.do("SPOP", "bbs"))
	fmt.Println(myredis.do("SRANDMEMBER", "bbs"))
	fmt.Println(myredis.do("SREM"), "bbs", "baidu.com")
	fmt.Println(myredis.do("SUNION", "bbs", "web"))
	fmt.Println(myredis.do("SUNIONSTORE", "bbs", "web", "total"))
*/
	// sort set
/*
	fmt.Println(myredis.do("ZADD", "page_rank", 10, "baidu.com"))
	fmt.Println(myredis.do("ZADD", "page_rank", 11, "360.com", 9, "bing.com"))
	fmt.Println(myredis.do("SADD", "page_rank", "worldpress.net"))

	fmt.Println(myredis.do("ZCARD", "page_rank"))
	fmt.Println(myredis.do("ZCOUNT", "page_rank", 10, 11))
	fmt.Println(myredis.do("ZINCRBY", "page_rank", "baidu.com", 2))
	fmt.Println(myredis.do("ZRANGE", "page_rank", 0, -1, "WITHSCORES"))

	fmt.Println(myredis.do("ZRANGEBYSCORE", "page_rank", "-inf", "+inf"))
	fmt.Println(myredis.do("ZRANK", "page_rank", "baidu.com"))
	fmt.Println(myredis.do("ZREM", "page_rank", "360.com"))
	fmt.Println(myredis.do("ZREMRANGEBYRANK","page_rank", 0, 1))
	fmt.Println(myredis.do("ZREMRANGEBYSCORE", "page_rank", 9, 10))
	fmt.Println(myredis.do("ZREVRANGE", "page_rank", 0, -1, "WITHSCORES"))
	fmt.Println(myredis.do("ZREVRANGEBYSCORE"), "page_rank", "+inf", "-inf")
	fmt.Println(myredis.do("ZREVRANK ", "page_rank", "bing.com"))
	fmt.Println(myredis.do("ZSCORE ", "page_rank", "bing.com"))

	fmt.Println(myredis.do("ZADD", "programmer", 2000, "peter", 4000, "jack", 200, "passi"))
	fmt.Println(myredis.do("ZADD", "manager", 5000, "herry", 6000, "mary", 200, "passi"))

	fmt.Println(myredis.do("ZUNIONSTORE ", "salary", 2, "programmer", "manager", "WEIGHTS", 1, 3))
	fmt.Println(myredis.do("ZINTERSTORE", "sum", 2, "programmer", "manager"))

*/



}