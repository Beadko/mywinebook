package cmd

import (
	"log"

	"github.com/Beadko/mywinebook/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See a list of wines you tried",
	Long:  `Get a full list wine entries that you have added to your diary `,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := data.GetWines(); err != nil {
			log.Println("Failed to get wines")
			return
		}
	},
}

func init() {
	wineCmd.AddCommand(listCmd)
}
