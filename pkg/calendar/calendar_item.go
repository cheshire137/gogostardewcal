package calendar

import "fmt"

type calendarItem struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Person string `json:"person"`
	day    int
	season string
}

func (i *calendarItem) Event() (Event, error) {
	if i.Type == "birthday" {
		return NewBirthday(i.Person, i.day, i.season), nil
	}
	if i.Type == "festival" {
		return NewFestival(i.Name, i.day, i.season), nil
	}
	if i.Type == "other" {
		return NewOtherEvent(i.Name, i.day, i.season), nil
	}
	return nil, fmt.Errorf("Error: invalid calendar item type '%s'", i.Type)
}
