package main

import (
	"context"
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configFile = flag.String("f", "configs/anime.yaml", "set config file which viper will loading.")
	flag.Parse()
	viper.AddConfigPath(*configFile)
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()
	gateway, err := New(ctx, *configFile)
	if err != nil {
		log.Fatal("new gateway err:", err)
	}
	err = gateway.Start()
	if err != nil {
		log.Fatal("gateway start err:", err)
	}
}
