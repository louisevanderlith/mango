package control

import (
	"strings"
)

type Menu map[string][]menuItem

var _activeLink string

func NewMenu(activeLink string) *Menu {
	_activeLink = activeLink

	menu := make(Menu)
	result := &menu
	result.SetActive(activeLink)

	return result
}

func (m *Menu) AddItem(link, text, iconClass string, children *Menu) {
	shortName, item := newItem("", link, text, iconClass, children)

	menu := *m
	menu[shortName] = append(menu[shortName], item)
	m = &menu
}

func (m *Menu) AddItemWithID(id, link, text, iconClass string, children *Menu) {
	shortName, item := newItem(id, link, text, iconClass, children)

	menu := *m
	menu[shortName] = append(menu[shortName], item)
	m = &menu
}

func (m *Menu) SetActive(link string) bool {
	foundActive := false

	for _, v := range *m {
		for _, item := range v {
			item.IsActive = item.Link == link

			if !foundActive && item.IsActive {
				foundActive = true
			}

			foundActiveChild := item.Children.SetActive(link)

			if foundActiveChild {
				item.IsActive = true
			}
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

func newItem(id, link, text, iconClass string, children *Menu) (shortName string, result menuItem) {
	shortName = getUniqueName(text)
	result = menuItem{
		ID:       id,
		Name:     text,
		Link:     link,
		Class:    iconClass,
		IsActive: false,
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
