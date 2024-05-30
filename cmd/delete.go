/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Beadko/mywinebook/data"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a note",
	Long:  `Delete the entire wine entry`,
	Run: func(cmd *cobra.Command, args []string) {
		DeleteNote()
	},
}

func init() {
	noteCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func DeleteNote() {
	namePromptContent := promptContent{
		"What wine info do you want to remove?",
		"Type the name you want to delete",
	}
	name := promptGetInput(namePromptContent)

	data.DeleteWine(name)
}
