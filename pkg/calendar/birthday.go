package calendar

import "fmt"

type Birthday struct {
	Person string
	day    int
	season string
}

var _ Event = (*Birthday)(nil)

func NewBirthday(person string, day int, season string) *Birthday {
	return &Birthday{Person: person, day: day, season: season}
}

func (b *Birthday) Day() int {
	return b.day
}

func (b *Birthday) Season() string {
	return b.season
}

func (b *Birthday) String() string {
	return fmt.Sprintf("%s's birthday", b.Person)
}
