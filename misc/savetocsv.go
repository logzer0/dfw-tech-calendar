package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func main() {

	getConfig("/repos/dfw-tech-calendar/keys.json")
	DumpEventInfo("bro.csv")
}

//DumpEventInfo retrieves all the meetup groups and writes them to a CSV file
func DumpEventInfo(filePath string) {
	var ok bool
	var v int
	var a [][]string
	v = 0

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Num", "Name", "Link", "Urlname", "Created", "City", "State", "Country", "LocalizedCountryName", "JoinMode", "Visibility", "ID", "Members", "OrganizerID", "Name", "Bio", "Who", "Timezone", "CategoryID", "CategoryName", "CategoryShortname", "CategorySortName", "Score", "Lat", "Lon", "Description"}

	writer.Write(header)

	for i := 0; i < 20; i++ {
		a, ok, v = constructCSVRows(config.Zip, i, v)
		writer.WriteAll(a)

		if ok {
			break
		}
	}
}

//constructCSVRows returns a slice of slice of strings that can be passed to the CSV writer
func constructCSVRows(zip string, offset int, v int) ([][]string, bool, int) {
	meetupTechGroupsURL := "https://api.meetup.com/find/groups?&sign=true&photo-host=public&zip=" + zip + "&page=1000&radius=50&key=" + config.MeetupKey
	url := meetupTechGroupsURL + "&offset=" + strconv.Itoa(offset)
	var myObj []Group
	var resp []byte
	var err error
	var returnVal [][]string

	if resp, err = Call(url); err != nil {
		log.Println("Expected:", nil, " Received: ", err)
	}

	err = json.Unmarshal(resp, &myObj)
	if err != nil {
		log.Println(err)
	}

	if len(myObj) > 0 {
		for _, group := range myObj {
			x := []string{strconv.Itoa(v), group.Name, group.Link, group.Urlname, strconv.FormatInt(group.Created, 10), group.City, group.State, group.Country, group.LocalizedCountryName, group.JoinMode, group.Visibility, strconv.Itoa(group.ID), strconv.Itoa(group.Members), strconv.Itoa(group.Organizer.ID), group.Organizer.Name, group.Organizer.Bio, group.Who, group.Timezone, strconv.Itoa(group.Category.ID), group.Category.Name, group.Category.Shortname, group.Category.SortName, strconv.FormatFloat(group.Score, 'f', -1, 64), strconv.FormatFloat(group.Lat, 'f', -1, 64), strconv.FormatFloat(group.Lon, 'f', -1, 64), group.Description}
			v++
			returnVal = append(returnVal, x)
		}
		return returnVal, false, v
	}
	return returnVal, true, v
}
