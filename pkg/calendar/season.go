package calendar

import "fmt"

type Season struct {
	Name string            `json:"name"`
	Days [][]*calendarItem `json:"days"`
}

func (s *Season) String() string {
	return s.Name
}

type calendarItem struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Person string `json:"person"`
}

func getSeasonByName(seasons []*Season, seasonName string) (*Season, error) {
	var season *Season
	for _, s := range seasons {
		if s.Name == seasonName {
			season = s
			break
		}
	}
	if season == nil {
		validSeasonNames := make([]string, len(seasons))
		for i, s := range seasons {
			validSeasonNames[i] = s.Name
		}
		return nil, fmt.Errorf("Error: invalid season '%s', choose from: %v", seasonName,
			validSeasonNames)
	}
	return season, nil
}
