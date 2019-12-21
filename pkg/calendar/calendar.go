package calendar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Calendar struct {
	currentDay    int
	currentSeason string
	seasons       []*Season
}

func NewCalendar(pathToCalendar string, day int, season string) (*Calendar, error) {
	seasons, err := loadCalendar(pathToCalendar)
	if err != nil {
		return nil, err
	}
	return &Calendar{currentDay: day, currentSeason: season, seasons: seasons}, nil
}

func (c *Calendar) String() string {
	return fmt.Sprintf("%s day %d", c.currentSeason, c.currentDay)
}

func loadCalendar(pathToCalendar string) ([]*Season, error) {
	file, err := readCalendarFile(pathToCalendar)
	if err != nil {
		return nil, err
	}
	seasons := make([]*Season, len(file.Seasons))
	for i, season := range file.Seasons {
		seasons[i] = &season
	}
	return seasons, nil
}

type calendarFile struct {
	Seasons []Season `json:"seasons"`
}

func readCalendarFile(pathToCalendar string) (*calendarFile, error) {
	bytes, err := ioutil.ReadFile(pathToCalendar)
	if err != nil {
		return nil, err
	}
	var data calendarFile
	json.Unmarshal(bytes, &data)
	return &data, nil
}
