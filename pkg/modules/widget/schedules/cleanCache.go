package widget_schedules

import (
	"os"
	"path/filepath"

	widget_service "github.com/Mateus-MS/Duo-Widget/modules/widget/service"
	"github.com/robfig/cron/v3"
)

func StartCleanCacheSchedule(widgetService widget_service.Iservice) {
	c := cron.New()

	c.AddFunc("0 0 * * *", func() {
		entries, err := os.ReadDir("./modules/widget/cache")
		if err != nil {
			return
		}

		for _, entry := range entries {
			if entry.IsDir() {
				err := os.RemoveAll(filepath.Join("./modules/widget/cache", entry.Name()))
				if err != nil {
					return
				}
			}
		}
	})

	c.Start()
}
