package main

import (
	"os"

	boxmate "github.com/Affix/BoxmateApi/pkg"
)

func main() {
	token, _ := boxmate.BoxmateLogin(os.Getenv("BOXMATE_EMAIL"), os.Getenv("BOXMATE_PASSWORD"))
	ev, _ := boxmate.GetTeamupEventsForDate(token.ApiKey, "2020-01-01")
	for _, i := range ev.Events {
		println(i.Name)
	}
}
