package calendar

type Birthday struct {
	Person string
	day    int
	season string
}

var _ Event = (*Birthday)(nil)

func NewBirthday(person string, day int, season string) *Birthday {
	return &Birthday{Person: person, day: day, season: season}
}
