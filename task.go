package main

import "github.com/jinzhu/gorm"

// Task base task struct
type Task struct {
	gorm.Model
	Depends       []*Task
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

	for _, dependTask := range data.Depend {
		dependTasks = append(dependTasks, RegisteredTasks[dependTask])
	}

	newTask := Task{
		Depends:       dependTasks,
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
