package uti

import "github.com/vicchass/utisli"

type Address struct {
	City           string
	Close_montreal bool
}

type Student struct {
	Firstname string
	Lastname  string
	Marks     []int
	Average   float64
	address   Address
}

var Stu1 = Student{Firstname: "james", Lastname: "block", Marks: []int{10, 12, 14, 15}, address: Address{City: "montreal", Close_montreal: true}}

var Stu2 = Student{Firstname: "bob", Lastname: "dylan", Marks: []int{8, 7, 14, 19}, address: Address{City: "laval", Close_montreal: true}}

var Stu3 = Student{Firstname: "mike", Lastname: "moralles", Marks: []int{8, 7, 6, 3}, address: Address{City: "boston", Close_montreal: false}}

var Stu4 = Student{Firstname: "elena", Lastname: "svitolina", Marks: []int{12, 15, 9, 8, 12, 18}, address: Address{City: "new-york", Close_montreal: false}}

var Stu5 = Student{Firstname: "janes", Lastname: "sawer", Marks: []int{4, 6, 9, 12, 15, 18}, address: Address{City: "montreal", Close_montreal: true}}

var Stu6 = Student{Firstname: "fransisco", Lastname: "palon", Marks: []int{17, 16, 19, 20, 2, 7, 3}, address: Address{City: "repentigny", Close_montreal: true}}

// check if a name is in the students , return an int of number of times the name exsits
func Check_name_exist(the_students []Student, firstname string) int {
	// return the slice of all firstname
	slice_firstname := Slice_string_first(the_students)
	// return number of first name
	num := utisli.Number_string_slice(slice_firstname, firstname)
	return num
}

// get a slice of strings of firstname
func Slice_string_first(all_students []Student) []string {
	var out_slice []string
	for _, v := range all_students {
		out_slice = append(out_slice, v.Firstname)
	}
	return out_slice

}
