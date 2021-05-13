package main

import (
	"context"
	"flag"
	"log"
)

func main() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
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
