package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type TaskActions interface {
	MarkAsDone()
	MarkAsUndone()
}

type Task struct {
	ID   int
	Name string
	Date string
	Done bool
}

type TaskDisplay interface {
	ShowTasks()
}

type TaskCollection struct {
	header []string
	tasks  []Task
}

func (t *TaskCollection) ShowTasks() {
	w := tabwriter.NewWriter(os.Stdout, 15, 20, 10, ' ', 0)
	// Print header
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", t.header[0], t.header[1], t.header[2], t.header[3])
	fmt.Fprintln(w, "----\t----\t----\t----")
	for _, task := range *&t.tasks {
		createdDateUTC, err := time.Parse(time.RFC3339, task.Date)
		if err != nil {
			fmt.Println("Error parsing UTC time")
			os.Exit(1)
		}

		diff := timediff.TimeDiff(createdDateUTC)
		fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", task.ID, task.Name, diff, task.Done)
	}
	w.Flush()
}

func (t *Task) MarkAsDone() {
	t.Done = true
}

func (t *Task) MarkAsUndone() {
	t.Done = false
}
