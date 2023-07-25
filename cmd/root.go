/*
Copyright © 2023 Antunes
*/
package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"log"
)

var surveyIconsConfig = func(icons *survey.IconSet) {
	icons.Question.Format = "cyan"
	icons.Question.Text = "[?]"
	icons.Help.Format = "blue"
	icons.Help.Text = "Help ->"
	icons.Error.Format = "yellow"
	icons.Error.Text = "Note ->"
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Cli to interact with ms-recommendation-plat-admin",
	Long:  "",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
