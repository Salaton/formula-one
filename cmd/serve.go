package cmd

import (
	"context"
	"log"

	"github.com/Salaton/formula-one/pkg/presentation"
	"github.com/Salaton/formula-one/pkg/presentation/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var Port int

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "This command will be used to run the server",
	Long:  `Serve is a command used to run the server`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			logger.NewLogger().Log(zerolog.FatalLevel, "subcommand `serve` only takes one argument")
		}
		initServer(Port)
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
	serveCmd.Flags().IntVarP(&Port, "port", "p", 0, "Specify port to run server on")
	serveCmd.MarkFlagRequired("port")
}

func initServer(port int) {
	ctx := context.Background()
	err := presentation.InitializeServer(ctx, port)
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}
}
