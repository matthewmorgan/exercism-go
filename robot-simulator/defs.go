package robot

import "fmt"

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

// Like an enum.  Assigns int values incremented with each assignment.
const (
	N Dir = iota
	E
	S
	W
)

type Dir int

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

type Action byte
// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}
