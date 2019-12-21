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

func (f *Festival) Day() int {
	return f.day
}

func (f *Festival) Season() string {
	return f.season
}

func (f *Festival) String() string {
	return f.Name
}
