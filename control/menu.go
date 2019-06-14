package control

import (
	"strings"
)

type Menu map[string]*menuItem

var _activeLink string

func NewMenu(activeLink string) *Menu {
	_activeLink = activeLink

	menu := make(Menu)
	result := &menu
	result.SetActive(activeLink)

	return result
}

func (m *Menu) AddItem(id, link, text, iconClass string, children *Menu) {
	shortName, item := newItem(id, link, text, iconClass, children)

	menu := *m
	menu[shortName] = item
	m = &menu
}

func (m *Menu) SetActive(link string) bool {
	foundActive := false

	for _, v := range *m {
		v.IsActive = v.Link == link

		if !foundActive && v.IsActive {
			foundActive = true
		}

		foundActiveChild := v.Children.SetActive(link)

		if foundActiveChild {
			v.IsActive = true
		}
	}

	return foundActive
}

type menuItem struct {
	ID       string
	Name     string
	Class    string
	Link     string
	IsActive bool
	Children *Menu `json:",omitempty"`
}

func newItem(id, link, text, iconClass string, children *Menu) (shortName string, result *menuItem) {
	shortName = getUniqueName(text)
	result = &menuItem{
		ID:       id,
		Name:     text,
		Link:     link,
		Class:    iconClass,
		IsActive: _activeLink == link,
	}

	if link == "#" {
		result.Link += shortName
	}

	if children != nil {
		result.Children = children
	} else {
		result.Children = NewMenu(link)
	}

	return shortName, result
}

func getUniqueName(raw string) string {
	return strings.ToLower(strings.Replace(raw, " ", "", -1))
}
