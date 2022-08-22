package main

import (
	"demo/service/price/cron/internal/config"
	"demo/service/price/cron/internal/schedule"
	"demo/service/price/cron/internal/svc"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/cron.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	cronRunner := cron.New()
	schedule.RegisterSchedule(cronRunner, ctx)
	c.MustSetUp()
	fmt.Println("Starting cron job")
	cronRunner.Start()

	select {}
}
