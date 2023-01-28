package app

import "github.com/bwmarrin/snowflake"

type UserId struct {
	Value snowflake.ID
}

type User struct {
	Id   UserId
	Name string
}

func CreateUser(name string) *User {
	return &User{Id: UserId{generateId()}, Name: name}
}
