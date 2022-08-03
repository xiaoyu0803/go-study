package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	cont, err := redis.Dial("tcp", "43.138.43.191:6379", redis.DialPassword("000415"))
	if err != nil {
		panic(err)
	}

	do, err := cont.Do("get", "a")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", do)
	defer cont.Close()
}
