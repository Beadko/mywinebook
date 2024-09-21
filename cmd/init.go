/*
Copyright © 2024 Beatrise Babra github.com/Beadko
*/
package cmd

import (
	"github.com/Beadko/mywinebook/db"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new mywinebook database and table",
	Long:  `Initialise a new mywinebook database and table`,
	Run: func(cmd *cobra.Command, args []string) {
		db.InitDB()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
