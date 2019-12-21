package calendar

type OtherEvent struct {
	Name   string
	day    int
	season string
}

var _ Event = (*OtherEvent)(nil)

func NewOtherEvent(name string, day int, season string) *OtherEvent {
	return &OtherEvent{Name: name, day: day, season: season}
}

func (e *OtherEvent) Day() int {
	return e.day
}

func (e *OtherEvent) Season() string {
	return e.season
}

func (e *OtherEvent) String() string {
	return e.Name
}
