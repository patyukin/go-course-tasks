package main

import (
	"fmt"

	"github.com/patyukin/go-course-tasks/hw-2/lib"
	"github.com/patyukin/go-course-tasks/hw-2/types"
	"github.com/patyukin/go-course-tasks/hw-2/uploader"
)

func printTable(data types.InputData) {
	for _, object := range data.Objects {
		fmt.Println("______________________")
		fmt.Printf("%s\t%s| Mean\n", object.Name, lib.Ternary(len(object.Name) > 8, "", "\t"))
		fmt.Println("______________________")

		grades := map[int][]int{}
		for _, result := range lib.Filter(data.Results, func(r types.Result) bool { return r.ObjectID == object.ID }) {
			student := lib.Filter(data.Students, func(s types.Student) bool { return s.ID == result.StudentID })[0]
			grades[student.Grade] = append(grades[student.Grade], result.Result)
		}

		totalResults := 0
		totalNumber := 0
		for grade, results := range grades {
			reduceMean := func(v int, r float64) float64 {
				return r + float64(v)
			}

			reduceTotal := func(v int, r int) int {
				return r + v
			}

			gradeMean := lib.Reduce(results, 0.0, reduceMean) / float64(len(results))
			fmt.Printf("%d grade \t| %.1f\n", grade, gradeMean)
			totalResults += lib.Reduce(results, 0, reduceTotal)
			totalNumber += len(results)
		}

		if totalNumber > 0 {
			mean := float64(totalResults) / float64(totalNumber)
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
