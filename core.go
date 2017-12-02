package main

import (
	"fmt"
)

func startTask(taskName string) error {
	task, err := getTask(Task{Name: taskName})
	if err != nil {
		return err
	}

	taskQuery, err := NewTaskQuery(*task)
	if err != nil {
		return err
	}
	queryId := GetFixedHash(6)
	taskQuery.id = queryId
	HaruhiTaskQuery[queryId] = &taskQuery

	for _, taskNode := range taskQuery.execQuery {
		go startTaskPipeLine(taskNode.task)
	}

	return nil
}

func startTaskPipeLine(task *Task) error {
	_, err := sendDirective(task, Directive{action: TASK_START})
	if err != nil {
		return fmt.Errorf("can not send start to task %v, error: %v", task.Name, err)
	}
	task.Status = TASK_START

	return nil
}
