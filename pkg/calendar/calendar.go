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

func (c *Calendar) NextDay() error {
	if c.CurrentDay == 28 {
		c.CurrentDay = 1
		nextSeason, err := getNextSeason(c.CurrentSeason, c.Seasons)
		if err != nil {
			return err
		}
		c.CurrentSeason = nextSeason
	} else {
		c.CurrentDay++
	}
	return nil
}

func (c *Calendar) PreviousDay() error {
	if c.CurrentDay == 1 {
		c.CurrentDay = 28
		previousSeason, err := getPreviousSeason(c.CurrentSeason, c.Seasons)
		if err != nil {
			return err
		}
		c.CurrentSeason = previousSeason
	} else {
		c.CurrentDay--
	}
	return nil
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
		lines = append(lines, " ")
		for _, event := range events {
			lines = append(lines, fmt.Sprintf("- %s", event))
			lines = append(lines, " ")
		}
	}

	return lines
}

func (c *Calendar) SeasonEmoji() string {
	return c.CurrentSeason.Emoji()
}

func (c *Calendar) DaySheet(lines ...string) string {
	width := 40
	height := 20
	var sb strings.Builder
	dateStr := fmt.Sprintf("%s %s %d", c.SeasonEmoji(), c.CurrentSeason, c.CurrentDay)
	dateStrRunes := []rune(dateStr)
	leftPadding := 3
	topPadding := 2
	dateLineRow := topPadding
	totalLines := len(lines)
	numerator := float64((height - dateLineRow) - (totalLines - 1))
	denominator := float64(2)
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
			} else if !topOrBottomRow && !firstOrLastColumn && (row != dateLineRow || contentIndex > len(dateStrRunes) || contentIndex < 0) {
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

func getNextSeason(season *Season, seasons []*Season) (*Season, error) {
	var lookingFor string
	if season.Name == "spring" {
		lookingFor = "summer"
	} else if season.Name == "summer" {
		lookingFor = "fall"
	} else if season.Name == "fall" {
		lookingFor = "winter"
	} else if season.Name == "winter" {
		lookingFor = "spring"
	} else {
		return nil, fmt.Errorf("Don't know next season after '%s'", season.Name)
	}
	for _, s := range seasons {
		if s.Name == lookingFor {
			return s, nil
		}
	}
	return nil, fmt.Errorf("Could not find season '%s'", lookingFor)
}

func getPreviousSeason(season *Season, seasons []*Season) (*Season, error) {
	var lookingFor string
	if season.Name == "spring" {
		lookingFor = "winter"
	} else if season.Name == "summer" {
		lookingFor = "spring"
	} else if season.Name == "fall" {
		lookingFor = "summer"
	} else if season.Name == "winter" {
		lookingFor = "fall"
	} else {
		return nil, fmt.Errorf("Don't know next season before '%s'", season.Name)
	}
	for _, s := range seasons {
		if s.Name == lookingFor {
			return s, nil
		}
	}
	return nil, fmt.Errorf("Could not find season '%s'", lookingFor)
}
