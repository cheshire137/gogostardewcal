package calendar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
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

func (c *Calendar) DaySheet(lines ...string) string {
	width := 40
	height := 20
	var sb strings.Builder
	dateStr := fmt.Sprintf("%s %d", c.CurrentSeason, c.CurrentDay)
	dateStrRunes := []rune(dateStr)
	leftPadding := 3
	topPadding := 2
	dateLineRow := topPadding
	contentStartRow := dateLineRow + 2
	lineIndex := 0
	totalLines := len(lines)
	var lineRunes []rune

	for row := 0; row < height; row++ {
		topOrBottomRow := row == 0 || row == height-1

		for column := 0; column < width; column++ {
			firstColumn := column == 0
			lastColumn := column == width-1
			firstOrLastColumn := firstColumn || lastColumn

			if topOrBottomRow && firstOrLastColumn {
				sb.WriteString("#")
			} else {
				if topOrBottomRow {
					sb.WriteString("-")
				}

				if firstOrLastColumn {
					sb.WriteString("|")
				}
			}

			leftContentIndex := column - leftPadding
			if lineIndex < totalLines {
				lineRunes = []rune(lines[lineIndex])
				lineIndex++
			}

			if row == dateLineRow && column >= leftPadding && leftContentIndex < len(dateStrRunes) {
				sb.WriteString(string(dateStrRunes[leftContentIndex]))
			} else if row == contentStartRow && column >= leftPadding && lineRunes != nil && leftContentIndex < len(lineRunes) {
				sb.WriteString(string(lineRunes[leftContentIndex]))
			} else if !topOrBottomRow && !firstOrLastColumn {
				sb.WriteString(" ")
			}

			if lastColumn {
				sb.WriteString("\n")
			}
		}
	}

	return sb.String()
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
