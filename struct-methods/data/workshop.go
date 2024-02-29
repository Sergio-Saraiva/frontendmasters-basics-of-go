package data

import "time"

type Workshop struct {
	Course
	Instructor
	Date time.Time
}

func (w Workshop) SingUp() bool {
	return true
}

func NewWorkshop(workshopName string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = workshopName
	w.Instructor = instructor
	return w
}
