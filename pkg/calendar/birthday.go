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
	lastChar := b.Person[len(b.Person)-1:]
	if lastChar == "s" {
		return fmt.Sprintf("%s' birthday", b.Person)
	}
	return fmt.Sprintf("%s's birthday", b.Person)
}
