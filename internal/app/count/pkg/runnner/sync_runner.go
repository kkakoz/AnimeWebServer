package runnner

import (
	"context"
	"github.com/robfig/cron"
	"log"
	"red-bean-anime-server/internal/app/count/domain"
)

func SyncCache(usecase domain.ICountUsecase) {
	c := cron.New()
	spec := "0 */2 * * * ?"
	err := c.AddFunc(spec, func() {
		usecase.UpdateIncr(context.Background())
	})
	if err != nil {
		log.Fatal("sync cache err:", err)
	}
	c.Start()
}
