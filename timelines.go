package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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
	fmt.Print(string(calData))
	var events []Event
	var calendar CalendarData

	yaml.Unmarshal([]byte(data), &events)
	yaml.Unmarshal([]byte(calData), &calendar)
	fmt.Printf("%+v\n", events)
	fmt.Printf("%+v\n", calendar)
}
