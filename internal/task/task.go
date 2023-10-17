package task

import "fmt"

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
	return fmt.Sprintf("Name: '%s' ID: %s Project: %s", task.Name, task.ID, task.Project)
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
	// fmt.Printf("selected task:'%s' id:'%s' project:'%s'\n", task.Name, task.ID, task.Project)
	return tasks[index], nil
}
