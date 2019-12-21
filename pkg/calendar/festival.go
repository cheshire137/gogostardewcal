package calendar

type Festival struct {
	Name   string
	day    int
	season string
}

var _ Event = (*Festival)(nil)

func NewFestival(name string, day int, season string) *Festival {
	return &Festival{Name: name, day: day, season: season}
}
