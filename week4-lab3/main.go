package main
import (
	"fmt"
	"errors"
)

type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}
func (s *Student) IsHonor() bool {
	return s.GPA >= 3.5
}
func (s *Student) Validate() error {
	if s.Name == ""{
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1 and 4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("GPA must be between 0.0 and 4.0")
	}
	return nil
}
func main(){
	// var st Student = Student{ID:"1",Name:"",Email:"butkot_s@su.ac.th",Year:3,GPA:3.8}
	// st := Student{ID:"1",Name:"Sippakorn",Email:"butkot_s@su.ac.th",Year:3,GPA:3.8}
	students := []Student{
		{ID:"1",Name:"Sippakorn",Email:"butkot_s@su.ac.th",Year:3,GPA:3.8},
		{ID:"2",Name:"moodeng",Email:"moodeng@su.ac.th",Year:3,GPA:3.0},
	}
	newStudent := Student{ID:"3",Name:"aaaa",Email:"aaaa@gmail.com",Year:4,GPA:3.2}
	students = append(students, newStudent)

	for i, student := range students {
	fmt.Printf("%d Honor %v\n",i,student.IsHonor())
	fmt.Printf("%d Validate %v\n",i,student.Validate())
	}
}