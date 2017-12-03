package main

import (
	"errors"
	"fmt"
)

type TaskNode struct {
	task  *Task
	input []*TaskNode
	state uint8 // state 0 1 2
}

type TaskQuery struct {
	id        string
	current   int
	nodes     map[string]*TaskNode
	execQuery []*TaskNode
}

var HaruhiTaskQuery map[string]*TaskQuery

func NewTaskQuery(task Task) (TaskQuery, error) {
	inputnode := TaskNode{
		task:  &task,
		input: make([]*TaskNode, 0),
	}

	query := TaskQuery{
		current:   0,
		nodes:     make(map[string]*TaskNode),
		execQuery: make([]*TaskNode, 0),
	}

	query.appendTask(&inputnode)
	err := query.topologySort()
	if err != nil {
		return TaskQuery{}, err
	}

	return query, nil
}

func (query *TaskQuery) appendTask(tasknode *TaskNode) {
	dependTasks := tasknode.task.Depends
	depTaskLen := len(dependTasks)
	if depTaskLen == 0 {
		query.nodes[tasknode.task.Name] = tasknode
		return
	}
	if depTaskLen == len(tasknode.input) {
		return
	}
	query.nodes[tasknode.task.Name] = tasknode

	// 当没有依赖或者依赖全部连接完时直接返回

	for _, dependTask := range dependTasks {
		var dependTasknode *TaskNode
		if node, exists := query.nodes[dependTask.Name]; exists {
			dependTasknode = node
		} else {
			dependTasknode = &TaskNode{
				task:  dependTask,
				input: make([]*TaskNode, 0),
			}
		}
		tasknode.input = append(tasknode.input, dependTasknode)
		query.appendTask(dependTasknode)
	}
}

func (query *TaskQuery) topologySort() error {
	nodesLen := len(query.nodes)
	for len(query.execQuery) != nodesLen {
		currentTaskName, err := query.findZeroInputNode()
		if err != nil {
			return err
		}

		currentNode := query.nodes[currentTaskName]
		query.delTopologyInput(currentNode)
		query.execQuery = append(query.execQuery, currentNode)
		delete(query.nodes, currentTaskName)
	}

	return nil
}

func (query *TaskQuery) findZeroInputNode() (string, error) {
	for name, node := range query.nodes {
		if len(node.input) == 0 {
			return name, nil
		}
	}

	return "", errors.New("find loop in task query")
}

func (query *TaskQuery) delTopologyInput(tasknode *TaskNode) {
	for _, node := range query.nodes {
		delIndex := -1
		for i, inputnode := range node.input {
			if inputnode == tasknode {
				delIndex = i
				break
			}
		}

		if delIndex != -1 {
			copy(node.input[delIndex:], node.input[delIndex+1:])
			node.input[len(node.input)-1] = nil
			node.input = node.input[:len(node.input)-1]
		}
	}
}

func (query *TaskQuery) run() {
	query.startCurrentTask()
}

func (query *TaskQuery) startCurrentTask() error {
	taskNode := query.execQuery[query.current]
	if len(taskNode.input) != 0 {
		input := make(map[string]interface{})
		for _, inputNode := range taskNode.input {
			output := inputNode.task.Output
			input[inputNode.task.Name] = output
		}

		_, err := sendDirective(*taskNode.task, Directive{action: TASK_INPUT, data: input})
		if err != nil {
			return err
		}
	}

	_, err := sendDirective(*taskNode.task, Directive{action: TASK_START})
	if err != nil {
		return err
	}

	return nil
}

func (query *TaskQuery) finish(data finishTaskData) error {
	taskNode, ok := query.nodes[data.taskName]
	if !ok {
		return fmt.Errorf("finish data taskName error %v", data.taskName)
	}

	taskNode.task.Output = data.output
	// finish TaskQuery
	if query.current == len(query.execQuery)-1 {
		fmt.Printf("/n Task Query Finish, Id: %v \n", query.id)
		// TODO: save in postgres
		delete(HaruhiTaskQuery, query.id)
		return nil
	}

	query.current += 1
	err := query.startCurrentTask()
	if err != nil {
		return err
	}

	return nil
}
