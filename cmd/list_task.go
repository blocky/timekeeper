package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/task"
)

var KeyValueMode bool

var listTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func init() {
	f := listTaskCmd.Flags()
	f.BoolVarP(&ListPretty, "pretty", "p", false, "list as pretty-fied JSON")
	f.BoolVarP(&KeyValueMode, "key-value", "k", false, "list as 'ID':'Task Name'")

	listCmd.AddCommand(listTaskCmd)
}

func listTasks() {
	var tasks task.Tasks
	err := json.Unmarshal(TasksJSON, &tasks)
	check(err)

	if ListPretty {

		bytes, err := json.MarshalIndent(tasks, " ", " ")
		check(err)

		fmt.Printf("%s\n", bytes)
	} else if KeyValueMode {

		tasks.PrintKeyValue()
	} else {

		bytes, err := json.Marshal(tasks)
		check(err)

		fmt.Printf("%s\n", bytes)
	}
}
