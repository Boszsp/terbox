package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewList(choices []string) List {
	return List{
		choices:  choices,
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (l List) Init() tea.Cmd {
	return nil
}

func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
		case "down", "j":
			if l.cursor < len(l.choices)-1 {
				l.cursor++
			}
		case "enter", " ":
			_, ok := l.selected[l.cursor]
			if ok {
				delete(l.selected, l.cursor)
			} else {
				l.selected[l.cursor] = struct{}{}
			}
		}
	}
	return l, nil
}

func (l List) View() string {
	s := ""
	for i, choice := range l.choices {
		cursor := " "
		if l.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := l.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	return s
}
