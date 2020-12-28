package main

import "fmt"

// defining the struct
type Student struct {
	Name, Id, Department string
}

// Main Function
func main() {
	s := Student{Name: "Ram", Id: "GO101",
		Department: "IT"}

	fmt.Println("Student Name: ", s.Name)
	fmt.Println("Student Department: ", s.Department)

	s.Department = "CSE"

	// Displaying the result
	fmt.Println("Student: ", s)
}