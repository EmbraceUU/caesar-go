package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rwClient *RWClient

func SetUp() {
	SimpleClient()
}

func SimpleClient() {
	cli := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		DB:          10,
		ReadTimeout: 30 * time.Second,
	}).WithContext(context.Background())

	if err := cli.Ping().Err(); err != nil {
		panic(err)
		return
	}

	rwClient = NewSimpleClient(cli)

}

func GetClient() (*RWClient, error) {
	if rwClient == nil {
		return nil, fmt.Errorf("rwClient is nil. ")
	}
	return rwClient, nil
}

func NewRWClient(sentinel *redis.Client, instances []*redis.Client) *RWClient {
	cli := new(RWClient)
	cli.sentinel = sentinel
	cli.instances = instances
	return cli
}

func NewSimpleClient(c *redis.Client) *RWClient {
	cli := new(RWClient)
	cli.client = c
	return cli
}

type RWClient struct {
	sentinel  *redis.Client
	instances []*redis.Client
	client    *redis.Client
}

func (cli *RWClient) R() *redis.Client {
	return cli.client
}

func (cli *RWClient) W() *redis.Client {
	return cli.client
}

func (cli *RWClient) RandomR() *redis.Client {
	return cli.client
}
