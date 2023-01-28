package app

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
	node = n
}

func generateId() snowflake.ID {
	return node.Generate()
}
