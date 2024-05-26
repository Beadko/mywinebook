/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Beadko/mywinebook/data"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the note",
	Long:  `Update the details in the note`,
	Run: func(cmd *cobra.Command, args []string) {
		UpdateNote()
	},
}

func init() {
	noteCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func UpdateNote() {
	namePromptContent := promptContent{
		"What wine would you like to change?",
		"Type the name you want to change",
	}
	name := promptGetInput(namePromptContent)

	newNamePromptContent := promptContent{
		"What would you like to change the name to?",
		"Type in the new name",
	}
	newName := promptGetInput(newNamePromptContent)

	data.AmendNote(newName, name)
}
