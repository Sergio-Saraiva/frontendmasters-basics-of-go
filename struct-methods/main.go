package main

import (
	"fmt"

	"frontendmaster.com/course/go-basics/struct-methods/data"
)

func main() {

	var courses [2]data.Signable

	max := data.Instructor{
		Id:        3,
		FirstName: "Max",
		LastName:  "Ortiz",
	}

	goCourse := data.Course{
		Id:         1,
		Name:       "Go Basics",
		Instructor: max,
	}

	wksp := data.NewWorkshop("Swift", max)

	courses[0] = goCourse
	courses[1] = wksp

	for _, course := range courses {
		fmt.Println(course.SingUp())
	}
}
