package timekeep

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Month int

func MakeMonth(m int) (Month, error) {
	if 0 >= m || m > 12 {
		return 0, fmt.Errorf("month is invalid: '%d'", m)
	}
	return Month(m), nil
}

func (m Month) String() string {
	if m < 10 {
		return fmt.Sprintf("0%d", m)
	}
	return fmt.Sprintf("%d", m)
}

type Day int

func MakeDay(d int) (Day, error) {
	if 0 >= d || d > 31 {
		return 0, fmt.Errorf("day is invalid: '%d'", d)
	}
	return Day(d), nil
}

func (d Day) String() string {
	if d < 10 {
		return fmt.Sprintf("0%d", d)
	}
	return fmt.Sprintf("%d", d)
}

type MilitaryTime string

func MakeMilitaryTime(s string) (MilitaryTime, error) {
	if len(s) != 4 {
		return "", fmt.Errorf("provide 4 digits (HourMinutes)")
	}
	return MilitaryTime(s), nil
}

type Task struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Project string `json:"projectId"`
}

type Date struct {
	Year  int
	Month Month
	Day   Day
	Start MilitaryTime
	Stop  MilitaryTime
}

func (t Date) String() string {
	return fmt.Sprintf("%d-%s-%s:%s-%s", t.Year, t.Month, t.Day, t.Start, t.Stop)
}

type Entry struct {
	Date    Date
	Task    Task
	Details string
}

func (e Entry) String() string {
	return fmt.Sprintf("## %s %s %s %s\n", e.Date, e.Task.Project, e.Task.ID, e.Details)
}

type Timekeeper struct {
	*bufio.Writer
	tasks []Task
}

func MakeTimekeeper(
	f *os.File,
	tasks []Task,
) Timekeeper {
	w := bufio.NewWriter(f)
	return Timekeeper{w, tasks}
}

func (tk *Timekeeper) AskTime() (Date, error) {
	year := 2023

	month, err := tk.AskMonth("what month? (integer)")
	if err != nil {
		return Date{}, err
	}

	day, err := tk.AskDay("what day? (integer)")
	if err != nil {
		return Date{}, err
	}

	start, err := tk.AskMilitaryTime("start hour? (HourMinutes)")
	if err != nil {
		return Date{}, err
	}

	stop, err := tk.AskMilitaryTime("stop hour? (HourMinutes)")
	if err != nil {
		return Date{}, err
	}

	date := Date{
		Year:  year,
		Month: month,
		Day:   day,
		Start: start,
		Stop:  stop,
	}
	return date, nil
}

func (tk *Timekeeper) AskTask() (Task, error) {
	tk.PrintTasks()

	num, err := tk.AskInt("what task? (id number)")
	if err != nil {
		return Task{}, err
	}

	if num < 0 || num > len(tk.tasks)-1 {
		return Task{}, fmt.Errorf("No task exists for number:%d", num)
	}
	task := tk.tasks[num]
	fmt.Printf("selected task:'%s' id:'%s' project:'%s'\n", task.Name, task.ID, task.Project)

	return task, nil
}

func (tk *Timekeeper) AskDetails() string {
	details := tk.Ask("what did you do?")
	fmt.Printf("details: %s\n", details)
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
		Date:    time,
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

func (tk *Timekeeper) AskInt(question string) (int, error) {
	answer := tk.Ask(question)
	i, err := strconv.Atoi(answer)
	if err != nil {
		return 0, fmt.Errorf("given non-int: '%s'", answer)
	}
	return i, nil
}

func (tk *Timekeeper) AskMonth(question string) (Month, error) {
	m, err := tk.AskInt(question)
	if err != nil {
		return 0, err
	}
	month, err := MakeMonth(m)
	if err != nil {
		return 0, err
	}
	return month, nil
}

func (tk *Timekeeper) AskDay(question string) (Day, error) {
	d, err := tk.AskInt(question)
	if err != nil {
		return 0, err
	}
	day, err := MakeDay(d)
	if err != nil {
		return 0, err
	}
	return day, nil
}

func (tk *Timekeeper) AskMilitaryTime(question string) (MilitaryTime, error) {
	mt, err := MakeMilitaryTime(tk.Ask(question))
	if err != nil {
		return "", err
	}
	return mt, nil
}

func (tk *Timekeeper) PrintTasks() {
	fmt.Printf("Tasks\n-----\n")
	for i, task := range tk.tasks {
		fmt.Printf("%d) %s\n", i, task.Name)
	}
	fmt.Printf("-----\n")
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
