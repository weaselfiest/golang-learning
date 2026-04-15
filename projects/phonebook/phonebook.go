package phonebook

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrContactExists   = errors.New("contact already exists")
	ErrContactNotFound = errors.New("contact not found")
)

type PhoneBook struct {
	contacts map[string]string
}

func NewPhoneBook() *PhoneBook {
	return &PhoneBook{contacts: make(map[string]string)}
}

func (p *PhoneBook) Add(name, phone string) error {
	if _, ok := p.contacts[name]; ok {
		return ErrContactExists
	}
	p.contacts[name] = phone
	return nil
}

func (p *PhoneBook) Find(name string) (string, error) {
	if phone, ok := p.contacts[name]; ok {
		return phone, nil
	}
	return "", ErrContactNotFound
}

func (p *PhoneBook) Delete(name string) error {
	if _, ok := p.contacts[name]; ok {
		delete(p.contacts, name)
		return nil
	}
	return ErrContactNotFound
}

func (p *PhoneBook) Print() {
	keys := make([]string, 0, len(p.contacts))
	for key := range p.contacts {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("Name: %s, Phone: %s\n", key, p.contacts[key])
	}
}

func (p *PhoneBook) List() map[string]string {
	newMap := make(map[string]string, len(p.contacts))
	for k, v := range p.contacts {
		newMap[k] = v
	}
	return newMap
}
