package structs

import (
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Year    int
	Address string
	Contact string
	Catchup bool
}

// (s *Student) is called Method Receiver
// this turns this function into a method for Student Struct
func (s *Student) ChangeName(newName string) {
	s.Name = newName
	fmt.Printf("Name inside function scope: %s\n", s.Name)
	fmt.Printf("pointer of s %p\n", s)
}
