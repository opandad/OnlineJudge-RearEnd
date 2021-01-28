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
package database

import (
	"OnlineJudge-RearEnd/configs"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

/*
@Title
ConnectRedisDatabase

@description
返回redis数据库链接，并且返回所需要的context

@param
nil

@return
db, context.Background, error (*redis.Client, context.Context, error)
*/
func ConnectRedisDatabase() (*redis.Client, context.Context, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.DATABASE_REDIS_SERVER_IP + ":" + configs.DATABASE_REDIS_SERVER_PORT,
		Password: configs.DATABASE_REDIS_PASSWORD,
		DB:       0,
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, context.Background(), errors.New("redis数据库连接失败，需要检查redis数据库是否能够正确连接！" + pong)
	} else {
		return rdb, context.Background(), nil
	}
}
