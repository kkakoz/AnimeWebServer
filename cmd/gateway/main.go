package main

import (
	"context"
	"flag"
	"github.com/spf13/viper"
	"log"
)

func main() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var configFile = flag.String("f", "gateway.yml", "set config file which viper will loading.")
	viper.AddConfigPath(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read conf failed:", err)
	}
	gateway, err := New(ctx, viper.GetViper())
	if err != nil {
		log.Fatal("new gateway err:", err)
	}
	err = gateway.Start()
	if err != nil {
		log.Fatal("gateway start err:", err)
	}
}
