package main

import (
	"fmt"
	"time"

	"github.com/logzer0/meetupGCal"
)

func main() {
	fmt.Println("Updating the calendar. Current time", time.Now())
	configFilePath := "keys.json"
	secretFilePath := "client_secret.json"
	meetupGCal.UpdateCalendar(configFilePath, secretFilePath)
}
