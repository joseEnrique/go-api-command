package main

// In the previous example we looked at
// [spawning external processes](spawning-processes). We
// do this when we need an external process accessible to
// a running Go process. Sometimes we just want to
// completely replace the current Go process with another
// (perhaps non-Go) one. To do this we'll use Go's
// implementation of the classic
// <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// function.

import (
	"fmt"

	"gopkg.in/redis.v5"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func getPID(key string) string {
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	return val
}

func setPID(key string, value string) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}

}

func main() {
	setPID("w", "aaaa")
	fmt.Println(getPID("w"))
}
