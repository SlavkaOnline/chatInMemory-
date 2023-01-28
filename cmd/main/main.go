package main

import (
	"chatInMemory/internal/app"
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func main() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}

	// Generate a snowflake ID.
	id := node.Generate()

	user := app.UserId{id}

	fmt.Println(user.Value)
}
