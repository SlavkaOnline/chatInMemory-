package app

import "github.com/bwmarrin/snowflake"

type RoomId struct {
	Value snowflake.ID
}

type Room struct {
	Id    RoomId
	Title string
}
