package auth

import (
	"github.com/google/wire"
	"time"
)

type ITokenGenerator interface {
	GenToken(id string) string
	GenTokenExpire(id string, expire time.Duration) string
}

var ProviderSet = wire.NewSet(NewJwtTokenVerifier, NewJwtTokenGen)