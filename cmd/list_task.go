package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/task"
)

var listTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func init() {
	listTaskCmd.Flags().BoolVarP(&ListPretty, "pretty", "p", false, "list as pretty-fied JSON")

	listCmd.AddCommand(listTaskCmd)
}

func listTasks() {
	var tasks task.Tasks
	err := json.Unmarshal(TasksJSON, &tasks)
	check(err)

	var bytes []byte
	if ListPretty {
		bytes, err = json.MarshalIndent(tasks, " ", " ")
	} else {
		bytes, err = json.Marshal(tasks)
	}
	check(err)

	fmt.Printf("%s\n", bytes)
}
