package cmd

import (
	"github.com/spf13/cobra"
)

var wineCmd = &cobra.Command{
	Use:   "wine",
	Short: "Your wine info here",
	Long:  `Your wine information here`,
}

func init() {
	rootCmd.AddCommand(wineCmd)
}
