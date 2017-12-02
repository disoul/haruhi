package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DirectiveState directive action types
type DirectiveState uint8

const (
	TASK_INPUT DirectiveState = iota
	TASK_OUTPUT
	TASK_START
	TASK_STOP
	TASK_RESTART
)

// Directive task control message
type Directive struct {
	action DirectiveState
	data   interface{}
}

type DirectiveResponse struct {
	status   int
	output   HaruhiOutput
	errorMsg string
}

func sendDirective(task Task, directive Directive) (DirectiveResponse, error) {
	var res DirectiveResponse
	body, err := json.Marshal(directive.data)
	if err != nil {
		return res, err
	}
	response, err := http.Post(task.DirectivePath, "application/json", bytes.NewReader(body))

	if err != nil {
		return res, err
	}

	if response.StatusCode != 200 {
		return res, fmt.Errorf("directive response unexpected: %v", response.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	err = json.Unmarshal(buf.Bytes(), &res)

	if err != nil {
		return res, err
	}

	return res, nil
}
