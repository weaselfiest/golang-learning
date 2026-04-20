package person

import (
	"fmt"
	"testing"
	"time"
)

func TestAge(t *testing.T) {
	tests := []struct {
		name     string
		person   *Person
		expected int
	}{
		{
			name:     "обычный возраст",
			person:   NewPerson("Alice", time.Date(1994, time.April, 15, 0, 0, 0, 0, time.UTC)),
			expected: 31,
		},
		{
			name:     "день рождения ещё не наступил в этом году",
			person:   NewPerson("Bob", time.Date(1990, time.December, 31, 0, 0, 0, 0, time.UTC)),
			expected: 35,
		},
		{
			name:     "день рождения сегодня",
			person:   NewPerson("Carol", time.Now().AddDate(-25, 0, 0)),
			expected: 25,
		},
		{
			name:     "новорождённый",
			person:   NewPerson("Baby", time.Now()),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.Age()
			if got != tt.expected {
				t.Errorf("Age() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		person   *Person
		expected string
	}{
		{
			name:     "стандартный формат",
			person:   NewPerson("Alice", time.Date(1994, time.April, 14, 0, 0, 0, 0, time.UTC)),
			expected: "Name: Alice, Age: 32",
		},
		{
			name:     "другое имя",
			person:   NewPerson("Bob", time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),
			expected: "Name: Bob, Age: 26",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.String()
			if got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		person   *Person
		expected string
	}{
		{
			name:     "стандартное приветствие",
			person:   NewPerson("Alice", time.Date(1994, time.April, 14, 0, 0, 0, 0, time.UTC)),
			expected: "Hi, my name is Alice and I am 32 years old!",
		},
		{
			name:     "другое имя",
			person:   NewPerson("Bob", time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),
			expected: "Hi, my name is Bob and I am 26 years old!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.Greet()
			if got != tt.expected {
				t.Errorf("Greet() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestStringerInterface(t *testing.T) {
	p := NewPerson("Alice", time.Date(1994, time.April, 14, 0, 0, 0, 0, time.UTC))
	expected := "Name: Alice, Age: 32"
	got := fmt.Sprintf("%v", p)
	if got != expected {
		t.Errorf("fmt.Sprintf(\"%%v\") = %q, want %q", got, expected)
	}
}
