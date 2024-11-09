package initialize

import (
	"context"
	"fiber/global"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"time"
)

// InitRedis 函数用于初始化Redis客户端并测试连接
func initRedis() {
	host := fmt.Sprintf("%s:%d", global.Conf.Redis.Host, global.Conf.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: global.Conf.Redis.Password, // 没有密码，默认值
		DB:       global.Conf.Redis.Database, // 默认DB 0
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		global.Logger.Panicf("initRedis client.Ping err:%v", err)

	}

	if err = redisotel.InstrumentTracing(client); err != nil {

		global.Logger.Panicf("initRedis client.Ping err: %v", err)
	}

	global.Logger.Info("Redis 初始化成功！")
	global.Redis = client
}
