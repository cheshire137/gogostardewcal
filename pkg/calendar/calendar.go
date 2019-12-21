package calendar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
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

func (c *Calendar) EventsSummary(events []Event) []string {
	lines := []string{}

	totalEvents := len(events)
	if totalEvents < 1 {
		lines = append(lines, "No events today")
	} else {
		plural := "s"
		if totalEvents == 1 {
			plural = ""
		}
		lines = append(lines, fmt.Sprintf("%d event%s today:", totalEvents, plural))
		for _, event := range events {
			lines = append(lines, fmt.Sprintf("- %s", event))
		}
	}

	return lines
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
	totalLines := len(lines)
	numerator := float64((height - 1) - dateLineRow)
	denominator := float64(totalLines)
	contentStartRow := int(math.Round(numerator / denominator))
	lineIndex := 0
	var lineRunes []rune

	for row := 0; row < height; row++ {
		topOrBottomRow := row == 0 || row == height-1
		wroteLine := false

		if lineIndex < totalLines {
			lineRunes = []rune(lines[lineIndex])
		} else {
			lineRunes = nil
		}

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

			contentIndex := column - leftPadding

			if row == dateLineRow && contentIndex >= 0 && contentIndex < len(dateStrRunes) {
				sb.WriteString(string(dateStrRunes[contentIndex]))
			} else if row >= contentStartRow && contentIndex >= 0 && lineRunes != nil && contentIndex < len(lineRunes) {
				sb.WriteString(string(lineRunes[contentIndex]))
				wroteLine = true
			} else if !topOrBottomRow && !firstOrLastColumn {
				sb.WriteString(" ")
			}

			if lastColumn {
				if wroteLine {
					lineIndex++
					wroteLine = false
				}
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
