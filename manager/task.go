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
