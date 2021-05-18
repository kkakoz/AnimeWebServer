package main

import (
	"context"
	"flag"
	"log"
	"red-bean-anime-server/internal/app/gateway"
)

func main() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, gateway.Gateway{}, struct {}{})
	var configFile = flag.String("f", "configs/gateway.yaml", "set config file which viper will loading.")
	gateway, err := New(ctx, *configFile)
	if err != nil {
		log.Fatal("new gateway err:", err)
	}
	err = gateway.Start()
	if err != nil {
		log.Fatal("gateway start err:", err)
	}
}
