package task

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Project string `json:"projectId"`
}

func MakeTask(
	name string,
	id string,
	project string,
) Task {
	return Task{
		Name:    name,
		ID:      id,
		Project: project,
	}
}

func (task Task) String() string {
	return fmt.Sprintf("Name:'%s' ID: %s Project: %s", task.Name, task.ID, task.Project)
}

func (task *Task) UnmarshalJSON(bytes []byte) error {
	type Alias Task
	err := json.Unmarshal(bytes, &struct{ *Alias }{Alias: (*Alias)(task)})
	if err != nil {
		return fmt.Errorf("could not unmarshal task: %s", err)
	}
	return nil
}

type Tasks []Task

func MakeTasks(tasks []Task) Tasks {
	return Tasks(tasks)
}

func (tasks Tasks) PrintTasks() {
	fmt.Printf("Tasks\n-----\n")
	for i, task := range tasks {
		fmt.Printf("%d) %s\n", i, task.Name)
	}
	fmt.Printf("-----\n")
}

func (tasks Tasks) SelectTask(index int) (Task, error) {
	if index < 0 || index > len(tasks)-1 {
		return Task{}, fmt.Errorf("No task exists for number:%d", index)
	}
	return tasks[index], nil
}
