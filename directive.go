package haruhi

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
