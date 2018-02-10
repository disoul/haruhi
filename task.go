package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

// Task base task struct
type Task struct {
	gorm.Model
	Depends       []*Task
	DependNames   pq.StringArray `gorm:"type:varchar(100)[]"`
	Type          string         `gorm:"size:50"`
	Name          string         `gorm:"size:50;unique"`
	DirectivePath string         `gorm:"size:100"`
	Status        DirectiveState `gorm:"-"`
	Output        HaruhiOutput   `gorm:"-"`
}

// RegisteredTasks tasks data in memory
var RegisteredTasks map[string]*Task

func registerTask(data RegisterData) error {
	dependTasks := make([]*Task, 0)
	dependTaskNames := make([]string, 0)
	for _, dependTask := range data.Depend {
		dependTasks = append(dependTasks, RegisteredTasks[dependTask])
		dependTaskNames = append(dependTaskNames, RegisteredTasks[dependTask].Name)
	}
	fmt.Printf("dep %v", dependTaskNames)

	newTask := Task{
		Depends:       dependTasks,
		DependNames:   dependTaskNames,
		Type:          data.Typename,
		Name:          data.Name,
		DirectivePath: data.Path,
	}

	task, err := saveTask(newTask)
	if err != nil {
		return err
	}

	RegisteredTasks[data.Name] = task

	return nil
}
