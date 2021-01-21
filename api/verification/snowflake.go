package verification

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

func Snowflake() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal("Snowflake fail: ", err)
	}
	return node.Generate().Int64()
}
