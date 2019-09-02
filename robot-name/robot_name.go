package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var usedNames = make(map[string]bool)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	if r.name == "" {
		r.Reset()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	usedNames[r.name] = false
	r.name = generateName()
	for usedNames[r.name] == true {
		r.name = generateName()
	}
	usedNames[r.name] = true
}

func generateName() string {
	return fmt.Sprintf("%s%s%03d", randomLetter(), randomLetter(), randomNumber(1000))
}

func randomLetter() string {
	return string(LETTERS[randomNumber(26)])
}

func randomNumber(max int) int {
	return rand.Intn(max)
}
