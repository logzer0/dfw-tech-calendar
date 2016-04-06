package main

import (
	"fmt"

	"github.com/logzer0/vegeta"
)

func main() {
	url := "https://api.meetup.com/find/groups?&sign=true&photo-host=public&zip=75063&category=34&page=20&key=629f7e19624e182a5c5117743"
	var resp []byte
	var err error
	handler := vegeta.NewHandler()
	if resp, err = handler.GetRequest(url); err != nil {
		fmt.Errorf("Expected:", nil, " Received: ", err)
	}
	if len(resp) < 1 {
		fmt.Errorf("Expected length to be greater than 0 got ", len(resp))
	}
	fmt.Println(string(resp))
}

type Group struct {
	Score                int     `json:"score"`
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Link                 string  `json:"link"`
	Urlname              string  `json:"urlname"`
	Description          string  `json:"description"`
	Created              int64   `json:"created"`
	City                 string  `json:"city"`
	Country              string  `json:"country"`
	LocalizedCountryName string  `json:"localized_country_name"`
	State                string  `json:"state"`
	JoinMode             string  `json:"join_mode"`
	Visibility           string  `json:"visibility"`
	Lat                  float64 `json:"lat"`
	Lon                  float64 `json:"lon"`
	Members              int     `json:"members"`
	Organizer            struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Bio   string `json:"bio"`
		Photo struct {
			ID          int    `json:"id"`
			HighresLink string `json:"highres_link"`
			PhotoLink   string `json:"photo_link"`
			ThumbLink   string `json:"thumb_link"`
		} `json:"photo"`
	} `json:"organizer"`
	Who        string `json:"who"`
	GroupPhoto struct {
		ID          int    `json:"id"`
		HighresLink string `json:"highres_link"`
		PhotoLink   string `json:"photo_link"`
		ThumbLink   string `json:"thumb_link"`
	} `json:"group_photo"`
	Timezone  string `json:"timezone"`
	NextEvent struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		YesRsvpCount int    `json:"yes_rsvp_count"`
		Time         int64  `json:"time"`
		UtcOffset    int    `json:"utc_offset"`
	} `json:"next_event"`
	Category struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Shortname string `json:"shortname"`
		SortName  string `json:"sort_name"`
	} `json:"category"`
	Photos []struct {
		ID          int    `json:"id"`
		HighresLink string `json:"highres_link"`
		PhotoLink   string `json:"photo_link"`
		ThumbLink   string `json:"thumb_link"`
	} `json:"photos"`
}

type Event struct {
	Created       int64  `json:"created"`
	Duration      int    `json:"duration"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	RsvpLimit     int    `json:"rsvp_limit"`
	Status        string `json:"status"`
	Time          int64  `json:"time"`
	Updated       int64  `json:"updated"`
	UtcOffset     int    `json:"utc_offset"`
	WaitlistCount int    `json:"waitlist_count"`
	YesRsvpCount  int    `json:"yes_rsvp_count"`
	Group         struct {
		Created  int64   `json:"created"`
		Name     string  `json:"name"`
		ID       int     `json:"id"`
		JoinMode string  `json:"join_mode"`
		Lat      float64 `json:"lat"`
		Lon      float64 `json:"lon"`
		Urlname  string  `json:"urlname"`
		Who      string  `json:"who"`
	} `json:"group"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Venue       struct {
		ID                   int     `json:"id"`
		Name                 string  `json:"name"`
		Lat                  float64 `json:"lat"`
		Lon                  float64 `json:"lon"`
		Repinned             bool    `json:"repinned"`
		Address1             string  `json:"address_1"`
		City                 string  `json:"city"`
		Country              string  `json:"country"`
		LocalizedCountryName string  `json:"localized_country_name"`
	} `json:"venue"`
	Visibility string `json:"visibility"`
}
