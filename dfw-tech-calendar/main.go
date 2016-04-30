package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/logzer0/meetupGCal"
	"github.com/rk/go-cron"
)

func init() {
	meetupGCal.ConfigFilePath = "/repos/dfw-tech-calendar/keys.json"
	cron.NewDailyJob(3, 0, 0, func(time.Time) {
		fmt.Println("Updating the calendar. Current time", time.Now())
		meetupGCal.UpdateCalendar()
	})
}

func main() {
	fmt.Println("Updating the calendar. Current time", time.Now())
	meetupGCal.UpdateCalendar()
	http.ListenAndServe(":9580", nil)
}
