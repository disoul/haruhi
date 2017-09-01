package main

import "testing"

func TestScheduler(t *testing.T) {
	var (
		taskA Task
		taskB Task
		taskC Task
		taskD Task
		taskE Task
		taskF Task
	)

	taskA = Task{
		Name:    "A",
		Depends: []*Task{&taskB},
	}

	taskB = Task{
		Name: "B",
	}

	taskC = Task{
		Name:    "C",
		Depends: []*Task{&taskA, &taskB},
	}

	taskD = Task{
		Name:    "D",
		Depends: []*Task{&taskA, &taskE},
	}

	taskE = Task{
		Name:    "E",
		Depends: []*Task{&taskF},
	}

	taskF = Task{
		Name:    "F",
		Depends: []*Task{&taskD},
	}

	tests := []struct {
		input  Task
		output string
	}{
		{taskC, "BAC"},
		{taskD, "error"},
	}

	for _, test := range tests {
		got := schedulerOutput(test.input)
		if got != test.output {
			t.Errorf(
				"input tast: %+v\nwant output: %s\noutput: %s",
				test.input,
				test.output,
				got,
			)
		}
	}
}

func schedulerOutput(task Task) string {
	query, err := NewTaskQuery(task)
	if err != nil {
		return "error"
	}

	output := ""
	for _, taskNode := range query.execQuery {
		output = output + taskNode.task.Name
	}

	return output
}
