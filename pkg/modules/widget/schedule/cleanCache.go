package widget_schedule

import (
	"os"
	"path/filepath"

	"github.com/robfig/cron/v3"
)

func StartCleanCacheSchedule() {
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
