package util

import runewidth "github.com/mattn/go-runewidth"

func SafeWidth(s string) int {
	return runewidth.StringWidth(s)
}
