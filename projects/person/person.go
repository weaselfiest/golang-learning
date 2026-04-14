package person

import (
	"fmt"
	"time"
)

type Person struct {
	name      string
	birthDate time.Time
}

func NewPerson(name string, date time.Time) *Person {
	return &Person{name: name, birthDate: date}
}

func (p *Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.Age())
}

func (p *Person) Age() int {
	now := time.Now()
	years := now.Year() - p.birthDate.Year()

	if now.Month() < p.birthDate.Month() ||
		(now.Month() == p.birthDate.Month() && now.Day() < p.birthDate.Day()) {
		years--
	}

	return years
}

func (p *Person) Greet() string {
	return fmt.Sprintf("Hi, my name is %s and I am %d years old!", p.name, p.Age())
}
