package main

import (
	"flag"
	"github.com/xiaoweiba-xiaoxiao/redis-cluser-cli/redis_command"
	"os"
	_ "github.com/xiaoweiba-xiaoxiao/goconfig/config"
)

var (
	cmdstr string
	filestr string 
    appname string
)


func init(){
	cfile := "redisconf.yml"
	appname = flag.CommandLine.Name()	
	flag.StringVar(&cmdstr,"cmd","","the requred option")
	flag.StringVar(&filestr,"f",cfile,"the unrequred option")
	flag.Parse()
}

func main(){
	if cmdstr == "" {
		flag.CommandLine.Usage()
		os.Exit(0)
	}
	redis_command.Run(filestr,cmdstr)	
}



