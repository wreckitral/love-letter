package cmd

import (
	"math/rand"
)

func GetShortLetter() (string, error) {
	n := rand.Intn(2)

	l := ShortLetters[n]

	return l, nil
}

func GetLongLetter() (string, error) {
	n := rand.Intn(2)

	l := LongLetters[n]

	return l, nil
}
