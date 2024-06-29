package main

import (
	"fmt"

	"github.com/patyukin/go-course-tasks/hw-2/cache"
	"github.com/patyukin/go-course-tasks/hw-2/types"
	"github.com/patyukin/go-course-tasks/hw-2/uploader"
)

func printTable(data types.InputData) {
	var studentCache cache.Cache[int, types.Student]
	var objectCache cache.Cache[int, types.Object]

	studentCache.Init()
	objectCache.Init()

	for _, student := range data.Students {
		studentCache.Set(student.ID, student)
	}

	for _, object := range data.Objects {
		objectCache.Set(object.ID, object)
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Student name | Grade\t  | Object\t| Result")
	fmt.Println("-----------------------------------------------")

	for _, result := range data.Results {
		student, _ := studentCache.Get(result.StudentID)
		object, _ := objectCache.Get(result.ObjectID)
		fmt.Printf("%s\t     | %d\t  | %s\t| %d\t\n", student.Name, student.Grade, object.Name, result.Result)
	}

	fmt.Println("-----------------------------------------------")
}

func main() {
	data := uploader.UploadHomeWorkData()
	printTable(data)
}
