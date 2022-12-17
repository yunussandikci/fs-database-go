package main

import (
	"fmt"
	"github.com/yunussandikci/kv-database-go/fsdatabase"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	db, dbErr := fsdatabase.New[[]Student]("test.db")
	if dbErr != nil {
		panic(dbErr)
	}

	students, readErr := db.Read()
	if readErr != nil {
		panic(readErr)
	}

	if students == nil {
		students = []Student{}
	}

	students = append(students, Student{
		Id:   len(students) + 1,
		Name: fmt.Sprintf("Student %d", len(students)+1),
	})

	for _, student := range students {
		fmt.Printf("Id:%d, Name:%s\n", student.Id, student.Name)
	}

	if saveErr := db.Write(students); saveErr != nil {
		return
	}
}
