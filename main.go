package main

import (
	"fmt"
	"playground/storage"

	"github.com/mediocregopher/radix.v2/pool"
)

func main() {
	// create Pool
	pool, err := pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		fmt.Errorf("pool: %v", err)
	}

	// create RedisStorage
	rs := &storage.RedisStorage{
		Pool: pool,
	}

	// test Set
	err = rs.Set("test1@example.com", "dfdfdf")
	if err != nil {
		fmt.Errorf("conn: %v", err)
	}

	// test Get
	data, err := rs.Get("test1@example.com")
	if err != nil {
		fmt.Errorf("conn: %v", err)
	}
	fmt.Println("-->", data)
}
