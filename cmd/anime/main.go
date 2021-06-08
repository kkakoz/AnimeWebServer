package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configFile = flag.String("f", "configs/anime.yaml", "set config file which viper will loading.")
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()
	app, err := NewApp(ctx, *configFile)
	if err != nil {
		log.Fatal("new app failed:", err)
	}
	err = app.Start()
	if err != nil {
		log.Fatal("run app failed:", err)
	}
}
