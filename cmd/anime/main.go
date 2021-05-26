package main

import (
	"context"
	"flag"
	"github.com/spf13/viper"
	"log"
)

func main() {
	var configFile = flag.String("f", "configs/anime.yaml", "set config file which viper will loading.")
	viper.AddConfigPath(*configFile)
	ctx := context.TODO()
	app, err := NewApp(ctx, *configFile)
	if err != nil {
		log.Fatal("run new failed:", err)
	}
	err = app.Start()
	if err != nil {
		log.Fatal("run app failed:", err)
	}
}
