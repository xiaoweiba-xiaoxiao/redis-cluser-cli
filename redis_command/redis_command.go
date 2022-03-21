package redis_command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
	cf "github.com/xiaoweiba-xiaoxiao/goconfig/config"
	"github.com/xiaoweiba-xiaoxiao/redis-cluser-cli/db_oprator"
)

func parse(conf string) ([]byte, error) {
	return cf.LoadYaml(conf)
}

func getHosts(configbyte []byte) []string {
	hosts := []string{}
	hgjons := gjson.GetBytes(configbyte, "redis.nodes").Array()
	for _, hgjon := range hgjons {
		host := hgjon.String()
		hosts = append(hosts, host)
	}
	if len(hosts) == 0 {
		panic("not find redis hosts or hosts is not array")
	}
	return hosts
}

func pasreCmd(cmd string) []string {
	return strings.Fields(cmd)
}

func run(file string, cmd string) {
	var err error
	configbyte, err := parse(file)
	if err != nil {
		fmt.Printf("pasre config file error:%v", err)
		return
	}
	hosts := getHosts(configbyte)
	client := db_oprator.NewClient(hosts)
	defer client.Close()
	cmdstring := pasreCmd(cmd)
	redisCmd := strings.ToLower(cmdstring[0])
	if !strings.HasSuffix(cmdstring[1], "*") {
		switch redisCmd {
		case "set":
			var str string
			str, err = client.Set(context.TODO(), cmdstring[1], cmdstring[2])
			if err == nil {
				fmt.Println(cmdstring[1], ":")
				fmt.Println(str)
			}
		case "get":
			var str string
			str, err = client.Get(context.TODO(), cmdstring[1])
			if err == nil {
				fmt.Println(cmdstring[1], ":")
				fmt.Println(str)
			}
		case "del":
			var integer int64
			integer, err = client.Del(context.TODO(), cmdstring[1])
			if err == nil {
				fmt.Println(cmdstring[1], ":")
				fmt.Println("integer:", integer)
			}
		case "lrange":
			var res = []string{}
			if len(cmdstring) == 4 {
				res, err = client.LRange(context.Background(), cmdstring[1:])
				if err == nil {
					fmt.Println(cmdstring[1], ":")
					fmt.Println(res)
				}
			} else {
				err = errors.New("the Usage: Lrange [key string] [start int64] [end int64]")
			}
		
		default:
			err = errors.New("unkown redis oprator!")
		}
		if err != nil {
			fmt.Printf("%v %v is failed:%v\n", cmdstring[0], cmdstring[1], err)
			return
		}
		fmt.Printf("%v %v is ok\n", cmdstring[0], cmdstring[1])
		return
	}
	ctx := context.Background()
	keys, err := client.Keys(ctx, cmdstring[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	switch redisCmd {
	case "del":
		for _, key := range keys {
			i, err := client.Del(context.Background(), key)
			fmt.Println(key, "is del begin")
			if err != nil {
				fmt.Println(key, "delete faild: ", err)
				continue
			}
			fmt.Println("result integer:", i)
			fmt.Println(key, "is del end")
		}
	case "keys":
		fmt.Println(keys)
	default:
		fmt.Printf("ERROR: unkown opreater of %s,cmd must be keys or del", cmdstring[1])
		return
	}
	fmt.Printf("%v %v is ok\n", cmdstring[0], cmdstring[1])
}

func Run(file string, cmd string) {
	run(file, cmd)
}
