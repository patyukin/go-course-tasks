package main

import (
	"fmt"

	"github.com/patyukin/go-course-tasks/hw-2/types"
	"github.com/patyukin/go-course-tasks/hw-2/uploader"
)

func printTable(data types.InputData) {
	students := make(map[int]types.Student)
	objects := make(map[int]types.Object)

	for _, student := range data.Students {
		students[student.ID] = student
	}

	for _, object := range data.Objects {
		objects[object.ID] = object
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Student name | Grade\t  | Object\t| Result")
	fmt.Println("-----------------------------------------------")

	for _, result := range data.Results {
		student := students[result.StudentID]
		object := objects[result.ObjectID]
		fmt.Printf("%s\t     | %d\t  | %s\t| %d\t\n", student.Name, student.Grade, object.Name, result.Result)
	}
}

func main() {
	data := uploader.UploadHomeWorkData()

	printTable(data)
}
