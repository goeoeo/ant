package redisutil

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type (

	//节点配置
	Node struct {
		Address  string //节点地址 127.0.0.1:6379
		Password string // 密码

		//redisPool 初始化参数
		MaxIdle     int //最大空闲连接数
		MaxActive   int //最大的激活连接数，同时最多有N个连接：0=不限制连接数
		IdleTimeout int //空闲连接等待时间，超过此时间后，空闲连接将被关闭：0=超时也不断开空闲连接

		Pool *redis.Pool
	}

	//统一接口
	Cli interface {
		Conn() redis.Conn
	}
)

//创建redisPool
func (this *Node) NewRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     this.MaxIdle,
		MaxActive:   this.MaxActive,
		IdleTimeout: time.Duration(this.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {

			c, err := redis.DialURL("redis://" + this.Address)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}

			//验证redis密码
			if this.Password != "" {
				if _, authErr := c.Do("AUTH", this.Password); authErr != nil {
					c.Close()
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

//输出连接
func (this *Node) Conn() redis.Conn {
	if this.Pool == nil {
		this.Pool = this.NewRedisPool()
	}

	return this.Pool.Get()
}
