package main

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
)

type Storage interface {
	Set(string)
	Get(string) string
}

type RedisStorage struct {
	Pool *pool.Pool
}

func main() {
	pool, err := pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		fmt.Errorf("pool: %v", err)
	}

	rs := RedisStorage{
		Pool: pool,
	}

	conn, err := rs.Pool.Get()
	if err != nil {
		fmt.Errorf("conn: %v", err)
	}
	defer conn.Close()

	resp := conn.Cmd("SET", "test1@example.com", "dfdfdf")
	if resp.Err != nil {
		fmt.Errorf("conn SET: %v")
	}

	foo, err := conn.Cmd("GET", "test1@example.com").Str()
	if err != nil {
		fmt.Errorf("conn GET: %v")
	}
	fmt.Println("-->", foo)
}
