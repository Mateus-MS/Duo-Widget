package widget_schedule

import (
	widgets_images "github.com/Mateus-MS/Duo-Widget/modules/widget/images"
	"github.com/robfig/cron/v3"
)

func GetActualMood() string {
	return mood
}

var mood string

func StartHourlyMoodRotator(moods *widgets_images.MoodRaw) {
	mood = moods.GetRandom()

	c := cron.New()
	c.AddFunc("0 * * * *", func() {
		mood = moods.GetRandom()
	})
	c.Start()
}
