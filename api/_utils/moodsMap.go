package utils

import "sync"

var moods map[string]string
var once sync.Once

func GetMood(mood string) string {

	once.Do(func() {
		moods = make(map[string]string)

		moods["greeting"] = "1"
		moods["passionate"] = "2"
		moods["sad"] = "3"
		moods["whistler"] = "4"
		moods["angry"] = "5"
		moods["chasing"] = "6"
		moods["splash"] = "7"
		moods["flirting"] = "8"
		moods["happy"] = "9"
		moods["delighted"] = "10"
		moods["splish"] = "11"
		moods["bored"] = "12"
		moods["struggling"] = "13"
	})

	return moods[mood]

}
