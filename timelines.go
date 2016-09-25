package main

import (
	"strings"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ParseDate (date string, calendar CalendarData) {
	split := strings.Split(date, " ")

	for i := 0; i < len(split); i++ {
		fmt.Printf("%+v\n", split[i])
	}
}

type Month struct {
	Name    string
	AltName string
	Days    int
}

type Date struct {
	Day       int
	Month     Month
	Year      int
	Reckoning string
}

type CalendarData struct {
	Name      string  `yaml:name`
	WeekCount int     `yaml:week_count`
	Months    []Month `yaml:months`
}

type Event struct {
	Summary     string `yaml:event`
	Date        string `yaml:date`
	Description string `yaml:description`
}

type Events struct {
	Events []Event
}

func main() {
	data, _ := ioutil.ReadFile("data/timeline-01.yml")
	calData, _ := ioutil.ReadFile("data/calendar-harpatos.yml")
	var events []Event
	var calendar CalendarData

	yaml.Unmarshal([]byte(data), &events)
	yaml.Unmarshal([]byte(calData), &calendar)

	ParseDate("1 Alturiak 1489 DR", calendar)
}
