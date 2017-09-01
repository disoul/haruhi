package haruhi

// TaskType task base type
type TaskType struct {
	Name string
}

// Task base task struct
type Task struct {
	Depends       []Task
	Type          TaskType
	Name          string
	DirectivePath string
}

// RegisteredTasks tasks data in memory
var RegisteredTasks map[string]Task

func registerTask(data RegisterData) error {
	dependTasks := make([]Task, 0)

	for _, dependTask := range data.Depend {
		dependTasks = append(dependTasks, RegisteredTasks[dependTask])
	}

	newTask := Task{
		Depends:       dependTasks,
		Type:          TaskType{data.Typename},
		Name:          data.Name,
		DirectivePath: data.Path,
	}

	err := saveTask(newTask)
	if err != nil {
		return err
	}

	RegisteredTasks[data.Name] = newTask

	return nil
}
