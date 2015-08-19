package main

import (
	"fmt"
	"gopkg.in/redis.v3"
	"net"
	"strconv"
)

const redisKey = "todo"

func GetRedisSlave() net.SRV {
	_, addrs, err := net.LookupSRV(*slaveService, *protocol, fmt.Sprintf("%s.%s", *framework, *domain))

	if err != nil {
		fmt.Println(err)
		return net.SRV{}
	}

	srvEntry := *(addrs[0])
	return srvEntry
}

func GetRedisMaster() net.SRV {
	_, addrs, err := net.LookupSRV(*masterService, *protocol, fmt.Sprintf("%s.%s", *framework, *domain))

	if err != nil {
		fmt.Println(err)
		return net.SRV{}
	}

	srvEntry := *(addrs[0])
	return srvEntry
}

func CreateRedisClient(host string, port string) *(redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Pong() {
	srvEntry := GetRedisSlave()
	client := CreateRedisClient(srvEntry.Target, strconv.Itoa(int(srvEntry.Port)))
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func RedisCreateTodo(name string) {
	srvEntry := GetRedisMaster()
	client := CreateRedisClient(srvEntry.Target, strconv.Itoa(int(srvEntry.Port)))
	cmd := client.RPush(redisKey, name)
	fmt.Println(cmd)
}

func RedisDeleteTodo(name string) {
	srvEntry := GetRedisMaster()
	client := CreateRedisClient(srvEntry.Target, strconv.Itoa(int(srvEntry.Port)))
	cmd := client.LRem(redisKey, 1, name)
	fmt.Println(cmd)
}

func RedisGetAllTodos() []string {
	srvEntry := GetRedisSlave()
	client := CreateRedisClient(srvEntry.Target, strconv.Itoa(int(srvEntry.Port)))
	cmd := client.LRange(redisKey, -100, 100)
	return cmd.Val()
}
