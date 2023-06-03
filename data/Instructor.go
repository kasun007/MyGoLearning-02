package data

import "fmt"

type Instructor struct { //struct is a collection of fields
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(name string, score int) Instructor { // factory function
	return Instructor{FirstName: name, Score: score}
}

func (i Instructor) Print() string { //method is a function with a special receiver argument

	return fmt.Sprintf("Id: %v\nFirst Name: %v\nLast Name: %v\nScore: %v\n", i.Id, i.FirstName, i.LastName, i.Score)
	//return fmt.Sprintf("Id: %v\nFirst Name: %v\nLast Name: %v\nScore: %v\n", i.Id, i.FirstName, i.LastName, i.Score)

}
