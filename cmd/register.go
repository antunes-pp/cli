/*
Copyright © 2023 Antunes
*/
package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/antunes-pp/cli/internal/adapter/question"
	"github.com/antunes-pp/cli/internal/core"
	"github.com/antunes-pp/cli/internal/core/user"
	"github.com/spf13/cobra"
)

var (
	qs = []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "Customer name:"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name:      "email",
			Prompt:    &survey.Input{Message: "Customer email:"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "squads",
			Prompt: &survey.MultiSelect{
				Message: "Choose squad to register this customer:",
				Options: core.SQUADS,
				Help:    "Choose at least one.",
			},
			Validate: survey.Required,
		},
		{
			Name: "confirm",
			Prompt: &survey.Confirm{
				Message: "Do you really want to register this new customer?",
			},
		},
	}
	isDev = false
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register user in squads",
	Long:  "",
	RunE:  runE,
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().BoolVarP(&isDev, "dev", "d", false, "Run it in the DEV environment")
}

func runE(_ *cobra.Command, _ []string) error {
	if isDev {
		fmt.Println("🚀 Running in DEV mode")
	} else {
		fmt.Println("🚀 Running in DEV mode")
	}

	answer := question.RegisterUser{}

	err := survey.Ask(qs, &answer, survey.WithIcons(surveyIconsConfig))

	if err != nil {
		return err
	}

	if !answer.Confirm {
		fmt.Println("❌ Cancelled!")

		return nil
	}

	registerUseCase := user.NewRegisterUserUserCase()

	if err := registerUseCase.Execute(answer, isDev); err != nil {
		return err
	}

	fmt.Printf("✅ %s has been successfully registered!\n", answer.GetName())

	return nil
}
