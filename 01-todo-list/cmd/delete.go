/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"todo/utils"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete --id <id>",
	Short: "Delete a task by ID",
	Long:  `Delete a task by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := cmd.Flags().GetInt("id")
		utils.DeleteTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Flags
	deleteCmd.Flags().IntP("id", "i", 0, "Task ID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
