package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type Teacher struct {
	Person
	Id string
}
type Student struct {
	Person
	Grade     float64
	Professor Teacher
}
type AcademicMember interface {
	Description()
}

func (x Teacher) Description() {
	fmt.Printf("Teacher: Id: %s, Name: %s and Age: %d", x.Id, x.Name, x.Age)
}
func (y Student) Description() {
	fmt.Printf("Name: %s, Age: %d, Teacher: %s and Grade: %f", y.Name, y.Age, y.Professor.Name, y.Grade)
}

func ShowDescription(obj AcademicMember) {
	obj.Description()
}
func main() {
	var teacher01 = Teacher{
		Person: Person{Name: "Teacher 01", Age: 34},
		Id:     "AD45G3",
	}
	var teacher02 = Teacher{
		Person: Person{Name: "Teacher 02", Age: 45},
		Id:     "DS45J1",
	}

	var student01 = Student{
		Person:    Person{Name: "Student 01", Age: 19},
		Grade:     8.12,
		Professor: teacher01,
	}
	var student02 = Student{
		Person:    Person{Name: "Student 02", Age: 18},
		Grade:     7.23,
		Professor: teacher02,
	}
	var student03 = Student{
		Person:    Person{Name: "Student 03", Age: 21},
		Grade:     9.76,
		Professor: teacher02,
	}

	ShowDescription(teacher01)
	ShowDescription(teacher02)

	ShowDescription(student01)
	ShowDescription(student02)
	ShowDescription(student03)

}
