package db_oprator

import (
	redis "github.com/go-redis/redis/v8"
)

type Client struct{
	rdb *redis.ClusterClient
	node []*nodeclient
}

type nodeclient struct{
	nrdb *redis.Client
}