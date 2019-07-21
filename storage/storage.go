package storage

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
)

type Storage interface {
	Set(string, string) error
	Get(string) (string, error)
}

type RedisStorage struct {
	Pool *pool.Pool
}

func (rs *RedisStorage) Set(key, value string) error {
	conn, err := rs.Pool.Get()
	if err != nil {
		fmt.Errorf("conn: %v", err)
		return err
	}
	defer conn.Close()

	resp := conn.Cmd("SET", key, value)
	if resp.Err != nil {
		fmt.Errorf("conn SET: %v")
		return err
	}

	return nil
}

func (rs *RedisStorage) Get(key string) (string, error) {
	conn, err := rs.Pool.Get()
	if err != nil {
		fmt.Errorf("conn: %v", err)
		return "", err
	}
	defer conn.Close()

	value, err := conn.Cmd("GET", "test1@example.com").Str()
	if err != nil {
		fmt.Errorf("conn GET: %v")
		return "", err
	}

	return value, nil
}
