package main

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
