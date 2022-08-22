package main

import (
	"demo/service/price/cronjob/internal/config"
	"demo/service/price/cronjob/internal/schedule"
	"demo/service/price/cronjob/internal/svc"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/cronjob.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	cronRunner := cron.New()
	schedule.RegisterSchedule(cronRunner, ctx)
	fmt.Println("Starting cron job")
	cronRunner.Start()

	select {}
}
