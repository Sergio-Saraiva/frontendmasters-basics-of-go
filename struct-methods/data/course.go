package data

import "fmt"

type Duration float64

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) SingUp() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("--- %v --- %v", c.Name, c.Instructor.FirstName)
}
