/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tranvlnh26/go-commit/internal"
)

var flags internal.Flags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-commit",
	Short: "✨ Make commits easier with Gemini 🚀",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		internal.Run(flags)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(
		&flags.Emoji,
		"emoji",
		false,
		"Add a gitmoji to the commit message")
	rootCmd.Flags().StringVar(
		&flags.Template,
		"template",
		"",
		"Specify a custom commit message template. e.g. --template \"Modified {GIT_BRANCH} | {COMMIT_MESSAGE}\"")
	rootCmd.Flags().StringVar(
		&flags.Language,
		"language",
		"english",
		"Specify the language to use for the commit message(default: english). e.g. --language english")
	rootCmd.Flags().StringVar(
		&flags.CommitType,
		"commit-type",
		"",
		"pecify the type of commit to generate. This will be used as the type in the commit message e.g. --commit-type feat")
}
