package cmd

import (
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/joypauls/notask/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Suffix = " Running"
		spin.Start()

		// inefficient, should just hydrate a config struct
		apiKey := viper.GetString("apiKey")
		databasedId := viper.GetString("databaseId")
		client := &http.Client{}

		qr, db := src.FetchBoard(client, databasedId, apiKey)
		spin.Stop()
		src.PrintBoard(qr, db)
	},
}

func init() {
	rootCmd.AddCommand(boardCmd)
}
