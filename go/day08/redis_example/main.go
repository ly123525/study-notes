package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println("conn redis fail", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("Hset", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(conn.Do("Hget", "books", "abc"))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
