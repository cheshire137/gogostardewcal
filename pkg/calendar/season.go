package calendar

type Season struct {
	Name string           `json:"name"`
	Days [][]calendarItem `json:"days"`
}

type calendarItem struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Person string `json:"person"`
}
