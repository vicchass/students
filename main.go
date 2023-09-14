package main

import (
	"fmt"
	"os"

	"github.com/vicchass/students/uti"
	"github.com/vicchass/utisli"
)

// global variable
var all_students []uti.Student

func main() {
	//asign all student to all_students variable
	all_students = append(all_students, uti.Stu1, uti.Stu2, uti.Stu3, uti.Stu4, uti.Stu5, uti.Stu6)

	//START LOGIC
	for {
		var choice string
		fmt.Println("what do you want to do?,a: print name and average,b: add a mark ,q : quit")
		fmt.Scan(&choice)
		switch choice {
		case "a":
			all_students = Update_average(all_students)
			Print_stu(all_students, 1)
		case "b":
			all_students = Add_mark(all_students)
		case "q":
			os.Exit(3)

		}

	}
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

// print student name , and/or average based on choise , 0 print just first name , 1 print firstname and avergae
func Print_stu(all_students []uti.Student, choice int) {
	switch choice {
	case 0:
		fmt.Println("here are the list of all students:")
		for k, v := range all_students {
			fmt.Println("the student number", k, " is called:", v.Firstname)
		}
	case 1:
		for _, v := range all_students {
			fmt.Println("the average of ", v.Firstname, "is ", v.Average)
		}
	}
}

// func add a mark to a student
func Add_mark(all_the_students []uti.Student) []uti.Student {
	var name_to_add string
	fmt.Println("Type name of student you want to add mark , to stop type : q")
	fmt.Scan(&name_to_add)
	if name_to_add == "q" {
		return all_the_students
	}
	num := uti.Check_name_exist(all_the_students, name_to_add)
	switch num {
	case 0:
		fmt.Println("couldnt find the firstname , try again")
	case 1:
		var new_mark int
		fmt.Println("type the mark to add to", name_to_add)
		fmt.Scan(&new_mark)
		for _, v := range all_the_students {
			if v.Firstname == name_to_add {
				v.Marks = append(v.Marks, new_mark)
			}
		}
		return all_the_students
		// TO DO CASE 2
	}

	return all_the_students
}
