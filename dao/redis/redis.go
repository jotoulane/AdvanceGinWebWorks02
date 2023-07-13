package redis

import (
	"fmt"
	"webWorks02/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			settings.Conf.RedisConfig.Host,
			settings.Conf.RedisConfig.Port,
		),
		Password: settings.Conf.RedisConfig.Password, // 密码
		DB:       settings.Conf.RedisConfig.DB,       // 数据库
		PoolSize: settings.Conf.RedisConfig.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	rdb.Close()
}
