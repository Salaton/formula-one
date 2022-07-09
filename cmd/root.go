package main

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "formula-one",
	Short: "A simple service to pull F1 data and run experiments using cutting edge technology",
	Long:  `formula-one is a CLI that is used to query F1 information from ergast, an experimental web service which provides a historical record of motor racing data for non-commercial purposes`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.formula-one.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
