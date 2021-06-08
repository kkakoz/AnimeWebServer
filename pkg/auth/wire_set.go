package auth

import (
	"github.com/google/wire"
)

var AuthSet = wire.NewSet(NewJwtTokenVerifier, NewJwtTokenGen)