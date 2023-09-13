package main

import (
	"fmt"

	"github.com/vicchass/students/uti"
	"github.com/vicchass/utisli"
)

// global variable
var all_students []uti.Student

func main() {
	//asign all student to all_students variable
	all_students = append(all_students, uti.Stu1, uti.Stu2, uti.Stu3, uti.Stu4, uti.Stu5, uti.Stu6)

	//START LOGIC
	all_students = Update_average(all_students)

}

// FUNCTIONS

// get average for all student
func Update_average(all_students []uti.Student) []uti.Student {
	// var out_stud []uti.Student
	for k, v := range all_students {
		if len(v.Marks) == 0 {
			fmt.Println(v.Firstname, " DONT HAVE ANY MARKS")
		} else {
			value := utisli.Average_slice_int(v.Marks)
			all_students[k].Average = value
		}
	}

	return all_students

}

// print student name and average
func Print_stu_av(all_students []uti.Student) {
	for _, v := range all_students {
		fmt.Println("the average of ", v.Firstname, "is ", v.Average)
	}
}


func Add_marks