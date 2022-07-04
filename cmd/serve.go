package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/Salaton/formula-one/config"
	"github.com/Salaton/formula-one/pkg/presentation"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This command will be used to run the server",
	Long:  `Serve is a command used to run the server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		config, err := config.LoadConfig()
		if err != nil {
			return err
		}

		err = presentation.InitializeServer(ctx, config.EnvConfig.Port)
		if err != nil {
			log.Fatalf("an error occured: %v", err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
