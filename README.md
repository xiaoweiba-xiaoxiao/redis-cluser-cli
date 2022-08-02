# redis-cluser-cli
一个操作redis-cluster的运维命令，目前提供运维需要的set,del,keys,lrange,get这五种操作，支持string,list两种数据类型操作，后期有机会就会完善。

使用方法
1.构建成redis-cluster-cli的可执行文件。
go build -o redis-cluster-cli

2.配置文件

cluster-config.yaml
#配置如下：

redis:
  nodes:
  - 127.0.0.1:7000
  - 127.0.0.1:7001
  - 127.0.0.1:7002

./redis-cluster-cli -f cluster-config.yaml -cmd "keys admin::*"
./redis-cluster-cli -f cluster-config.yaml -cmd "del admin::*"
./redis-cluster-cli -f cluster-config.yaml -cmd "get admin::0357"






