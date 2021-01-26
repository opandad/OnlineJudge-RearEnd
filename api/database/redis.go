package database

import (
	"OnlineJudge-RearEnd/configs"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

/*
	@Title
	api/database/redis.go

	@Description
	redis数据库的连接以及测试数据库是否连接成功

	@Func List

	| func name             | develop  | unit test |

	|----------------------------------------------|

	| ConnectRedisDatabase  |    ok    |    ok	   |

	@config database path => ~/configs/redis.go
*/
func ConnectRedisDatabase() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.DATABASE_REDIS_SERVER_IP + ":" + configs.DATABASE_REDIS_SERVER_PORT,
		Password: configs.DATABASE_REDIS_PASSWORD,
		DB:       0,
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis ping error:", err)
	} else {
		fmt.Println("Redis ping result: ", pong)
	}

	return rdb
}
