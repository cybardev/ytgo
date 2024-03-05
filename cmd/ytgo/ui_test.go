package main

import (
	"testing"
)

func TestGetShortcut(t *testing.T) {
	var ctrlMap map[int]rune = map[int]rune{
		0:  rune(48),
		1:  rune(48 + 1),
		9:  rune(48 + 9),
		10: rune(87 + 10),
		36: rune(87 + 36),
		37: 0,
		-1: 0,
	}
	for i, r := range ctrlMap {
		symbol := getShortcut(i)
		if symbol != r {
			t.Errorf("Expected shortcut to be %d, got %d", r, symbol)
		}
	}
}
