package db_oprator

import (
	"context"
	"fmt"
	"strconv"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func NewClient(hosts []string)(*Client){
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: hosts,
		DialTimeout: 60 * time.Second,
	})
	client := &Client{rdb: rdb}
	_,err := client.rdb.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	return client
}

func (c *Client)String() string{
	return fmt.Sprintf("%#v",c.rdb)
}

func (c *Client)Del(ctx context.Context,key string)(int64,error){
	intcmd := c.rdb.Del(ctx,key)
	return intcmd.Result()
}

func (c *Client)Get(ctx context.Context,key string)(string,error){
	stringcmd := c.rdb.Get(ctx,key)
	return stringcmd.Result()
}

func (c *Client)LRange(ctx context.Context,args []string)([]string,error){
	startint,err := strconv.ParseInt(args[1],10,64)
	if err != nil {
		return nil,err
	}
	endint,err := strconv.ParseInt(args[2],10,64)
	if err != nil {
		return nil,err
	}
	stringcmd := c.rdb.LRange(ctx,args[0],startint,endint)
	return stringcmd.Result()
}

func (c *Client)LKeys(ctx context.Context,partten string)([]string,error){
	keys := []string{}
	err := c.rdb.ForEachMaster(ctx,func(ctx context.Context, rdb *redis.Client) error{
		 r := rdb.Keys(ctx,partten)
		 keys = append(keys,r.Val()...)
		 return r.Err()
	})
	return keys,err
}

func (c *Client)Set(ctx context.Context,key string,value interface{})(string,error){
	statuscmd := c.rdb.Set(ctx,key,value,0)
	return statuscmd.Result()
}



func (c *Client)Do(ctx context.Context,args ...interface{})(interface{},error){
	rdbcmd := c.rdb.Do(ctx,args...)
	return rdbcmd.Result()
}

func (c *Client)Keys(ctx context.Context,partten string)([]string,error){
	var keys []string
	err := c.rdb.ForEachMaster(ctx,func(ctx context.Context, rdb *redis.Client) error {
		iter := rdb.Scan(ctx, 0, partten, 2000).Iterator()
		for iter.Next(ctx) {
			keys = append(keys,iter.Val())
		}
		return iter.Err()
	})
	return keys,err
}

func (c Client)Close(){
	c.rdb.Close()
}

