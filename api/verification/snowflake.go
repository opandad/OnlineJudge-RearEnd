package verification

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

/*
	@Title
	api/snowflake.go

	@Description
	雪花算法，生成不重复id

	@Func List（这个需打开函数检查）

	| func name            | develop  | unit test |

	|---------------------------------------------|

	| Snowflake            |    no    |    no	  |
*/
func Snowflake() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal("Snowflake fail: ", err)
	}

	return node.Generate().String()
}
