package timekeep

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Task struct {
	Name    string
	ID      string
	project string
}

type Time struct {
	Date  string
	Start string
	Stop  string
}

func (t Time) String() string {
	return fmt.Sprintf("%s:%s-%s", t.Date, t.Start, t.Stop)
}

type Entry struct {
	Time    Time
	Task    Task
	Details string
}

func (e Entry) String() string {
	return fmt.Sprintf("## %s %s %s %s\n", e.Time, e.Task.project, e.Task.ID, e.Details)
}

var tasks = map[int]Task{
	1:  {"Meetings", "6109a1f50a335f28d1599ef3", "5f47d5879d6dc04fbfedcdab"},
	2:  {"ZPR - EKS K8 v1.23", "6446789ab958c4718390013a", "5f91ec0fb1d41c38c2d6719b"},
	3:  {"ZPR - DevOps Alpha", "64625befb2a18e3578754eb7", "5f91ec0fb1d41c38c2d6719b"},
	4:  {"ZPR - 10k Requests", "63aee243a3e51625ef082042", "5f91ec0fb1d41c38c2d6719b"},
	5:  {"IPR 3", "64625b8487e7196c5c4c3df6", "5f91ec0fb1d41c38c2d6719b"},
	6:  {"Infinity Pool of Crises", "5f91ef2eb1d41c38c2d6932f", "5f91ec0fb1d41c38c2d6719b"},
	7:  {"Dugnad", "5f91eca7b1d41c38c2d67837", "5f91ec0fb1d41c38c2d6719b"},
	8:  {"ZPR - Update Mocks", "643405d2cb8fab2f0d10f6b5", "5f91ec0fb1d41c38c2d6719b"},
	9:  {"ZPR - Strong Verify", "63aee21ca3e51625ef081f39", "5f91ec0fb1d41c38c2d6719b"},
	10: {"White Paper 0.3", "643405bbcb8fab2f0d10f430", "5f91ec0fb1d41c38c2d6719b"},
	11: {"COGS", "63f766a08c5c0806c7bb015b", "5f91ec0fb1d41c38c2d6719b"},
	12: {"Security Policy Implementation", "635008cb000c0f0aca116dad", "5f91ec0fb1d41c38c2d6719b"},
	13: {"Market Research", "6304cf552f286f53f966fbcc", "5f91ec0fb1d41c38c2d6719b"},
	14: {"ZPR - Detached Cacheit", "62c72c4710ace715d5ceda7f", "5f91ec0fb1d41c38c2d6719b"},
	15: {"SIA Proposal", "643405635cbbcc37df575eb0", "5f91ec0fb1d41c38c2d6719b"},
	16: {"Grant Admin", "61f15163c32016098b513e86", "5f91ec0fb1d41c38c2d6719b"},
	17: {"NSDI Submission", "643406065cbbcc37df577158", "5f91ec0fb1d41c38c2d6719b"},
}

func printTasks() {
	var list = make([]int, len(tasks))
	var i int

	for key, _ := range tasks {
		list[i] = key
		i++
	}
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })

	for _, i := range list {
		task := tasks[i]
		fmt.Printf("%d) %s\n", i, task.Name)
	}
}

type Timekeeper struct {
	*bufio.Writer
}

func MakeTimekeeper(f *os.File) Timekeeper {
	w := bufio.NewWriter(f)
	return Timekeeper{w}
}

func (tk *Timekeeper) AskTime() (Time, error) {
	day := tk.Ask("what day? (YYYY-MM-DD)")
	start := tk.Ask("start hour? (HHHH)")
	stop := tk.Ask("stop hour? (HHHH)")
	return Time{Date: day, Start: start, Stop: stop}, nil
}

func (tk *Timekeeper) AskTask() (Task, error) {
	printTasks()

	num := tk.Ask("what task? (id number)")
	taskNum, err := strconv.Atoi(num)
	if err != nil {
		return Task{}, fmt.Errorf("Given non int string: %s", num)
	}

	task, exists := tasks[taskNum]
	if !exists {
		return Task{}, fmt.Errorf("No task exists for number:%d", taskNum)
	}
	return task, nil
}

func (tk *Timekeeper) AskDetails() string {
	details := tk.Ask("what did you do?")
	return details
}

func (tk *Timekeeper) MakeEntry() (Entry, error) {
	time, err := tk.AskTime()
	if err != nil {
		return Entry{}, err
	}
	task, err := tk.AskTask()
	if err != nil {
		return Entry{}, err
	}
	details := tk.AskDetails()

	return Entry{
		Time:    time,
		Task:    task,
		Details: details,
	}, nil
}

func (tk *Timekeeper) WriteEntry(entry Entry) error {
	_, err := tk.WriteString(entry.String())
	if err != nil {
		return err
	}
	tk.Flush()
	return nil
}

func (tk *Timekeeper) Ask(question string) string {
	answer := prompt(question)
	return answer
}

func prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
