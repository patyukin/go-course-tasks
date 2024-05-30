package main

import (
	"fmt"

	"github.com/patyukin/go-course-tasks/hw-2/lib"
	"github.com/patyukin/go-course-tasks/hw-2/types"
	"github.com/patyukin/go-course-tasks/hw-2/uploader"
)

func printTable(data types.InputData) {
	excellentStudents := lib.Filter(data.Students, func(s types.Student) bool {
		studentResults := lib.Filter(data.Results, func(r types.Result) bool {
			return r.StudentID == s.ID
		})

		for _, result := range studentResults {
			if result.Result != 5 {
				return false
			}
		}
		return len(studentResults) > 0
	})

	fmt.Println("Круглые отличники:")
	for _, student := range excellentStudents {
		fmt.Printf("ID: %d, Name: %s, Grade: %d\n", student.ID, student.Name, student.Grade)
	}
}

func main() {
	data := uploader.UploadHomeWorkData()
	printTable(data)
}
