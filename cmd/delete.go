package cmd

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Beadko/mywinebook/data"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the wine",
	Long:  `Delete the wine entry`,
	Run: func(cmd *cobra.Command, args []string) {
		var ID string
		err := survey.AskOne(&survey.Input{
			Message: "Enter the ID of the wine you want to delete",
			Help:    "The ID should be a string",
		}, &ID)
		if err != nil {
			log.Fatalf("Failed to get the wine: %v", err)
		}
		err = data.DeleteWine(ID)
		if err != nil {
			log.Fatalf("Failed to delete wine: %v", err)
		}
	},
}

func init() {
	wineCmd.AddCommand(deleteCmd)
}
