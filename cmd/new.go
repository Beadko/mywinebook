/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Beadko/mywinebook/data"
	"github.com/Beadko/mywinebook/internal/wine"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

type promptContent struct {
	errorMsg string
	label    string
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Created a new note",
	Long:  `Creates a new note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

func init() {
	noteCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %s\n", result)
	return result
}

func createNewNote() {
	namePromptContent := promptContent{
		"Please add a wine Name",
		"Add a wine name",
	}
	name := promptGetInput(namePromptContent)

	typePromptContent := promptContent{
		fmt.Sprintf("What is the type of %s?", name),
		"Add a wine type?",
	}
	t := promptGetSelect(typePromptContent)
	tInt, err := strconv.Atoi(t)
	if err != nil {
		return
	}
	data.AddWine(name, tInt)
}

func promptGetSelect(pc promptContent) string {
	var result string
	var err error

	prompt := promptui.Select{
		Label: pc.label,
		Items: wine.WineType{},
	}

	_, result, err = prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %s\n", result)
	return result
}
