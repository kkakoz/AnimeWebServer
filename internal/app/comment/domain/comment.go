package domain

import "red-bean-anime-server/pkg/db/mysqlx"

type Comment struct {
	mysqlx.Model
	UserId int64 `json:"user_id"`
	Content string `json:"content"`

}