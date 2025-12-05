package libs

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

type RedisCfg struct {
	Addr   string `json:"addr"`
	Passwd string `json:"passwd"`
}

type RedisClusterCfg struct {
	Addrs  []string `json:"addrs"`
	Passwd string   `json:"passwd"`
}

func ConnRedis(c *RedisCfg) redis.Cmdable {
	ret := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Passwd,
	})
	if err := ret.Ping(context.Background()).Err(); err != nil {
		log.Print(err.Error())
		return nil
	}
	return ret
}

func ConnRedisCluster(c *RedisClusterCfg) redis.Cmdable {
	ret := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    c.Addrs,
		Password: c.Passwd,
	})
	if err := ret.Ping(context.Background()).Err(); err != nil {
		log.Print(err.Error())
		return nil
	}
	return ret
}
