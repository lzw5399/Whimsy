package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("start request redis...")
	conn, err := redis.Dial("tcp", "123.207.67.80:6379,password=redis@pwd,poolsize=5")
	if err != nil{
		fmt.Println("request redis failed",err)
		return
	}
	defer conn.Close()

}
