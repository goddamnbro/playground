package storage

import (
	"github.com/mediocregopher/radix.v2/pool"
)

type Storage interface {
	Set(string)
	Get(string) string
}

type RedisStorage struct {
	Pool *pool.Pool
}

func (rs *RedisStorage) Set(key string) {
	return
}

func (rs *RedisStorage) Get(key string) string {
	return ""
}
