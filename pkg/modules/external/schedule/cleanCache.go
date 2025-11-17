package external_schedule

import (
	external_service "github.com/Mateus-MS/Duo-Widget/modules/external/service"
	"github.com/robfig/cron/v3"
)

func StartCleanCacheSchedule(externalService external_service.Iservice) {
	c := cron.New()

	c.AddFunc("0 0 * * *", func() {
		externalService.Reset()
	})

	c.Start()
}
