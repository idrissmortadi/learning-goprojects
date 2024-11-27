/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"todo/utils"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName, _ := cmd.Flags().GetString("task")

		task := utils.Task{
			Name: taskName,
			Date: time.Now().UTC().Format(time.RFC3339), // Current time
		}

		fmt.Println("Adding task: ", task)
		utils.SaveTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.PersistentFlags().String("task", "", "Task name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
