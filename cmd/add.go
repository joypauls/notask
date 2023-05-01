package cmd

import (
	"net/http"

	"github.com/joypauls/notask/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("NOTION_API_KEY")
		databasedId := viper.GetString("NOTION_DATABASE_ID")
		// fmt.Println(args[0])
		newPage := src.PageRequest{
			Parent:     src.Parent{DatabaseId: databasedId},
			Properties: src.Properties{Name: src.Name{Title: []src.Title{{src.Text{Content: args[0]}}}}},
		}
		client := &http.Client{}
		src.InsertPage(client, newPage, databasedId, apiKey)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
