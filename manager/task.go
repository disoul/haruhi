package manager

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

// RegisteredTask tasks data in memory
var RegisteredTasks map[string]Task

func registerTask(name string, typename string, path string, depends []string) {
	dependTasks := make([]Task, 0)

	for _, dependTask := range depends {
		dependTasks = append(dependTasks, RegisteredTasks[dependTask])
	}

	// TODO: save in mongo

	RegisteredTasks[name] = Task{
		dependTasks,
		TaskType{typename},
		name,
		path,
	}
}
