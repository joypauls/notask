package cmd

import (
	"net/http"

	"github.com/joypauls/notask/src"
	"github.com/spf13/cobra"
)

var getBoardCmd = &cobra.Command{
	Use:     "getBoard",
	Aliases: []string{"board"},
	Short:   "Print the version number of Hugo",
	Long:    `All software has versions. This is Hugo's`,
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// inefficient, should just hydrate a config struct
		apiKey := src.ReadDotEnvFile("NOTION_API_KEY")
		databasedId := src.ReadDotEnvFile("NOTION_DATABASE_ID")
		client := &http.Client{}
		src.PrintBoard(client, databasedId, apiKey)
	},
}

func init() {
	rootCmd.AddCommand(getBoardCmd)
}
