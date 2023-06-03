package data

import "time"

type Workshop struct {
	Course     Course
	Name       string
	Date       time.Time
	Instructor Instructor
}

func (Workshop) SignUp() bool {
	return true
}

func NewWorkshop(name string, intructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name

	w.Instructor = intructor
	return w

}
