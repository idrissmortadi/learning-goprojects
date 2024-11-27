/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"todo/utils"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete --id <id>",
	Short: "Complete a task by ID",
	Long:  `Complete a task by ID; set the column completed to 'true'`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Println("Couldn't parse ID")
			os.Exit(1)
		}
		utils.CompleteTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	completeCmd.Flags().IntP("id", "i", -1, "ID of completed task")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
