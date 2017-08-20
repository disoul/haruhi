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

func registerTask(data RegisterData) {
	dependTasks := make([]Task, 0)

	for _, dependTask := range data.Depend {
		dependTasks = append(dependTasks, RegisteredTasks[dependTask])
	}

	// TODO: save in mongo

	RegisteredTasks[data.Name] = Task{
		Depends:       dependTasks,
		Type:          TaskType{data.Typename},
		Name:          data.Name,
		DirectivePath: data.Path,
	}
}
