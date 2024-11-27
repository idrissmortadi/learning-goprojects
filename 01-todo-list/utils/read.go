package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func GetReader(filepath string) (*csv.Reader, *os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, nil, err
	}

	r := csv.NewReader(f)

	return r, f, nil

}

func LoadRecords(filepath string) (TaskCollection, error) {
	r, f, err := GetReader(filepath)
	defer f.Close()

	rawRecords, err := r.ReadAll()
	if err != nil {
		return TaskCollection{}, fmt.Errorf("failed to read records from csvReader: %w", err)
	}

	// Initialize the header
	header := rawRecords[0]

	// Initialize tasks as a slice and not a pointer to a slice
	tasks := []Task{}

	for i, rawRecord := range rawRecords {
		if i == 0 {
			continue
		}

		if len(rawRecord) != 4 {
			return TaskCollection{}, fmt.Errorf("invalid record format")
		}

		var (
			id   int
			name string
			date string
			done bool
		)

		// Convert strings to their respective types
		id, err = strconv.Atoi(rawRecord[0])
		if err != nil {
			return TaskCollection{}, fmt.Errorf("failed to convert ID: %w", err)
		}

		name = rawRecord[1]
		date = rawRecord[2]

		doneStr := rawRecord[3]
		if doneStr == "true" || doneStr == "false" {
			done, err = strconv.ParseBool(doneStr)
			if err != nil {
				return TaskCollection{}, fmt.Errorf("failed to convert Done: %w", err)
			}
		} else {
			return TaskCollection{}, fmt.Errorf("invalid Done value: %s", doneStr)
		}

		tasks = append(tasks, Task{
			ID:   id,
			Name: name,
			Date: date,
			Done: done,
		})
	}

	// Return the initialized TaskCollection
	return TaskCollection{header: header, tasks: tasks}, nil
}
