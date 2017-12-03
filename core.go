package main

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

	go taskQuery.run()

	return nil
}
