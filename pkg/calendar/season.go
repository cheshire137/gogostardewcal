package calendar

import "fmt"

type Season struct {
	Name string            `json:"name"`
	Days [][]*calendarItem `json:"days"`
}

func (s *Season) GetEvents(day int) ([]Event, error) {
	if day < 1 || day > len(s.Days)+1 {
		return nil, fmt.Errorf("Error: invalid day %d", day)
	}
	items := s.Days[day-1]
	events := make([]Event, len(items))
	for i, item := range items {
		item.day = day
		item.season = s.Name
		event, err := item.Event()
		if err != nil {
			return nil, err
		}
		events[i] = event
	}
	return events, nil
}

func (s *Season) String() string {
	return s.Name
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
