package calendar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Calendar struct {
	CurrentDay    int
	CurrentSeason *Season
	Seasons       []*Season `json:"seasons"`
}

func NewCalendar(pathToCalendar string, day int, seasonName string) (*Calendar, error) {
	calendar, err := loadCalendar(pathToCalendar)
	if err != nil {
		return nil, err
	}

	season, err := getSeasonByName(calendar.Seasons, seasonName)
	if err != nil {
		return nil, err
	}

	calendar.CurrentDay = day
	calendar.CurrentSeason = season
	return calendar, nil
}

func (c *Calendar) CurrentEvents() ([]Event, error) {
	return c.CurrentSeason.GetEvents(c.CurrentDay)
}

func (c *Calendar) String() string {
	return fmt.Sprintf("%s day %d", c.CurrentSeason, c.CurrentDay)
}

func loadCalendar(pathToCalendar string) (*Calendar, error) {
	bytes, err := ioutil.ReadFile(pathToCalendar)
	if err != nil {
		return nil, err
	}
	var calendar Calendar
	json.Unmarshal(bytes, &calendar)
	return &calendar, nil
}
