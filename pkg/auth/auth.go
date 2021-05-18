package auth

import "time"

type ITokenGenerator interface {
	GenToken(id string) string
	GenTokenExpire(id string, expire time.Duration) string
}

