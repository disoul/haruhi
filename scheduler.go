package main

import (
	"errors"
)

type TaskNode struct {
	task  Task
	input []*TaskNode
	state uint8 // state 0 1 2
}

type TaskQuery struct {
	nodes     map[string]*TaskNode
	execQuery []*TaskNode
}

func NewTaskQuery(task Task) (TaskQuery, error) {
	inputnode := TaskNode{
		task:  task,
		input: make([]*TaskNode, 0),
	}

	query := TaskQuery{
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

func (query TaskQuery) appendTask(tasknode *TaskNode) {
	query.nodes[tasknode.task.Name] = tasknode
	dependTasks := tasknode.task.Depends

	// 当没有依赖或者依赖全部连接完时直接返回
	if depTaskLen := len(dependTasks); depTaskLen == 0 || depTaskLen == len(tasknode.input) {
		return
	}

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

func (query TaskQuery) topologySort() error {
	for len(query.execQuery) != len(query.nodes) {
		currentTaskName, err := query.findZeroInputNode()
		if err != nil {
			return err
		}

		currentNode := query.nodes[currentTaskName]
		query.delTopologyInput(currentNode)
		query.execQuery = append(query.execQuery, currentNode)
	}

	return nil
}

func (query TaskQuery) findZeroInputNode() (string, error) {
	for name, node := range query.nodes {
		if len(node.input) == 0 {
			return name, nil
		}
	}

	return "", errors.New("find loop in task query")
}

func (query TaskQuery) delTopologyInput(tasknode *TaskNode) {
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
