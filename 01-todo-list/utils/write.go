package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func GetWriter(filepath string) (*csv.Writer, *os.File, error) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file for writing")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, nil, err
	}

	w := csv.NewWriter(f)

	return w, f, nil
}

func SaveTask(task Task) {
	r, fr, err := GetReader("tasks.csv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawRecords, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fr.Close()

	var idInt int

	// Take last id
	if len(rawRecords) > 1 {
		id := rawRecords[len(rawRecords)-1][0]
		idInt, err = strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		idInt = 0
	}

	w, fw, err := GetWriter("tasks.csv")
	defer fw.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w.Write([]string{strconv.Itoa(idInt + 1), task.Name, task.Date, "false"})

	w.Flush()
}

func DeleteTask(id int) {
	r, fr, err := GetReader("tasks.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawRecords, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, rawRecord := range rawRecords {
		if rawRecord[0] == strconv.Itoa(id) {
			rawRecords = append(rawRecords[:i], rawRecords[i+1:]...)
			break
		}
	}

	fr.Close()

	// Overwrite file
	f, err := os.Create("tasks.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := csv.NewWriter(f)
	w.WriteAll(rawRecords)
	w.Flush()
}

func CompleteTask(id int) {
	r, fr, err := GetReader("tasks.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawRecords, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, rawRecord := range rawRecords {
		if rawRecord[0] == strconv.Itoa(id) {
			rawRecord[3] = "true"
		}
	}
	fr.Close()

	// Overwrite file
	f, err := os.Create("tasks.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := csv.NewWriter(f)
	w.WriteAll(rawRecords)
	w.Flush()
}
