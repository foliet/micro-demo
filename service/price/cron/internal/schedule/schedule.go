package schedule

import (
	"context"
	"demo/service/price/cron/internal/logic"
	"demo/service/price/cron/internal/svc"
	"github.com/robfig/cron"
	"log"
)

func MustAddJob(runner *cron.Cron, spec string, job cron.Job) {
	err := runner.AddJob(spec, job)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterSchedule(runner *cron.Cron, serverCtx *svc.ServiceContext) {
	MustAddJob(runner, "0 0 * * * ?", logic.NewShopeeScarperLogic(context.Background(), serverCtx))
}
