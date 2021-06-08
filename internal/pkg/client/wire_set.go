package client

import "github.com/google/wire"

var ClientSet = wire.NewSet(NewAnimeClient, NewCountClient, NewUserClient)
