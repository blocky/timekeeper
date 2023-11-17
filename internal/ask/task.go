package ask

import (
	"fmt"

	"github.com/blocky/timekeeper/internal/task"
)

func AskTask(tasks task.Tasks) (task.Task, error) {
	tasks.PrintTasks()

	index, err := AskInt("what task? (index number)")
	if err != nil {
		return task.Task{}, err
	}
	task, err := tasks.GetTask(index)
	if err != nil {
		return task, err
	}
	fmt.Printf("selected task: %+v\n", task)

	return task, nil
}
