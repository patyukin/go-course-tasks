package main

import (
	"fmt"
	"github.com/patyukin/go-course-tasks/hw-2/lib"
	"github.com/patyukin/go-course-tasks/hw-2/types"
	"github.com/patyukin/go-course-tasks/hw-2/uploader"
)

const longestName = 8

func printTable(data types.InputData) {
	students := make(map[int]types.Student)
	objects := make(map[int]types.Object)
	results := make(map[int]map[int][]int)

	for _, student := range data.Students {
		students[student.ID] = student
	}

	for _, object := range data.Objects {
		objects[object.ID] = object
	}

	for _, result := range data.Results {
		student := students[result.StudentID]
		if _, ok := results[result.ObjectID]; !ok {
			results[result.ObjectID] = make(map[int][]int)
		}

		results[result.ObjectID][student.Grade] = append(results[result.ObjectID][student.Grade], result.Result)
	}

	for object, grades := range results {
		fmt.Println("______________________")
		fmt.Printf("%s\t%s| Mean\n", objects[object].Name, lib.Ternary(len(objects[object].Name) > 8, "", "\t"))
		fmt.Println("______________________")

		total := 0
		count := 0

		for grade, value := range grades {
			gradeTotal := 0
			for _, result := range value {
				gradeTotal += result
			}

			gradeMean := float64(gradeTotal) / float64(len(value))
			fmt.Printf("%d grade \t| %.1f\n", grade, gradeMean)
			total += gradeTotal
			count += len(value)
		}

		if count > 0 {
			mean := float64(total) / float64(count)
			fmt.Println("______________________")
			fmt.Printf("mean\t\t| %.1f\n", mean)
			fmt.Println("======================")
		}
	}
}

func main() {
	data := uploader.UploadHomeWorkData()
	printTable(data)
}
