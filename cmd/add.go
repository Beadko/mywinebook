package cmd

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Beadko/mywinebook/data"
	"github.com/Beadko/mywinebook/internal/wine"
	"github.com/guregu/null/v5/zero"
	"github.com/spf13/cobra"
)

var addWineCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new wine to the database",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		var wineType zero.Int
		var country zero.Int
		var score zero.Int

		// Prompt for wine name
		err := survey.AskOne(&survey.Input{
			Message: "Enter the name of the wine:",
		}, &name)
		if err != nil {
			log.Printf("Failed to get the wine name: %v", err)
		}

		// Prompt for wine type
		err = survey.AskOne(&survey.Input{
			Message: "Enter the type of the wine:",
			Help:    "Type should be an integer",
		}, &wineType)
		if err != nil {
			log.Printf("Failed to get the wine type: %v", err)
		}

		// Prompt for country
		err = survey.AskOne(&survey.Input{
			Message: "Enter the country of the wine:",
			Help:    "Country should be an integer",
		}, &country)
		if err != nil {
			log.Printf("Failed to get the country: %v", err)
		}

		// Prompt for score
		err = survey.AskOne(&survey.Input{
			Message: "How would you score the wine?",
			Help:    "Score should be an integer",
		}, &score)
		if err != nil {
			log.Printf("Failed to get the score: %v", err)
		}

		// Add the wine to the database
		err = data.AddWine(wine.Wine{
			Name:      name,
			TypeID:    wineType,
			CountryID: country,
			Score:     score,
		})
		if err != nil {
			log.Printf("Failed to add wine: %v", err)
		}
	},
}

func init() {
	// Add command to root
	rootCmd.AddCommand(addWineCmd)
}
