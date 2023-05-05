package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

// // This is overwritten at compile time with build flags with the current tag
// // See build step in Makefile to get a sense of what happens
// var version = "v0.0.0"

// // Location to check for config override file
// // const configFile = ".scry/config.yaml"

// const titleText = "Notask CLI - a tool for Notion"
// const helpText = `Usage:
//   scry                   (Basic)
//   scry [flags] <path>    (Optional)

// Path:
//   <path> is a single optional argument that scry will try to resolve
//   to a valid starting directory. Default is the current directory.

// Flags:`

// func formatUsageText() string {
// 	return fmt.Sprintf("%s\n\n%s", titleText, helpText)
// }

// func printUsageText() {
// 	fmt.Fprintln(os.Stderr, formatUsageText())
// 	flag.PrintDefaults()
// }

// func parseArgs(args []string, c *app.Config) {
// 	if len(args) == 0 {
// 		// c.InitDir = fst.NewPath("")
// 	} else if len(args) == 1 {
// 		parsed, err := fp.Abs(args[0])
// 		if err != nil {
// 			log.Fatalf("Couldn't parse the path: %s", args[0])
// 		}
// 		fi, err := os.Stat(parsed)
// 		if os.IsNotExist(err) {
// 			log.Fatalf("No such file or directory: %s", args[0])
// 		} else if !fi.IsDir() {
// 			// parsed = fp.Dir(parsed)
// 		}
// 		// c.InitDir = fst.NewPath(parsed)
// 		fmt.Printf("Arg: %s\n", c.InitDir)
// 	} else {
// 		log.Fatal("Too many arguments supplied - zero(0) or one(1) required")
// 	}
// }

func FetchDatabase(client *http.Client, id string, key string) (Database, error) {
	requestURL := fmt.Sprintf("https://api.notion.com/v1/databases/%s", id)
	requestAuthValue := fmt.Sprintf("Bearer %s", key)
	request, _ := http.NewRequest("GET", requestURL, nil)
	request.Header.Add("Authorization", requestAuthValue)
	request.Header.Add("Notion-Version", "2022-06-28")

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var db Database
	err = json.Unmarshal(body, &db)
	// if err := json.Unmarshal(body, &db); err != nil { // Parse []byte to the go struct pointer
	// 	fmt.Println("Can not unmarshal JSON")
	// }
	return db, err
}

func FetchPages(client *http.Client, filter FilterPages, id string, key string) (QueryResult, error) {
	// build request
	requestURL := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", id)
	requestAuthValue := fmt.Sprintf("Bearer %s", key)
	bodyJson, _ := json.Marshal(filter)
	request, _ := http.NewRequest("POST", requestURL, bytes.NewBuffer(bodyJson))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", requestAuthValue)
	request.Header.Add("Notion-Version", "2022-06-28")

	// do request
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	// read result
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
	var qr QueryResult
	err = json.Unmarshal(body, &qr)
	// if err := json.Unmarshal(body, &qr); err != nil { // Parse []byte to the go struct pointer
	// 	fmt.Println("Can not unmarshal JSON")
	// }
	return qr, err
}

func InsertPage(client *http.Client, newPage PageRequest, id string, key string) {
	// build request
	requestURL := "https://api.notion.com/v1/pages"
	requestAuthValue := fmt.Sprintf("Bearer %s", key)
	bodyJson, _ := json.Marshal(newPage)
	request, _ := http.NewRequest("POST", requestURL, bytes.NewBuffer(bodyJson))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", requestAuthValue)
	request.Header.Add("Notion-Version", "2022-06-28")

	// do request
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	// read result
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(string(body))
	// var qr QueryResult
	// err = json.Unmarshal(body, &qr)
	// return qr, err
}

func FetchBoard(client *http.Client, id string, key string) (QueryResult, Database) {

	//Create user struct which need to post.
	filter := FilterPages{
		Filter: Filter{Property: "State", Status: Status{"Not started"}},
	}
	qr, _ := FetchPages(client, filter, id, key)
	// fetchedPages := qr.Results
	db, _ := FetchDatabase(client, id, key)
	// dbName := db.Title[0].Text.Content

	return qr, db
}

func PrintBoard(qr QueryResult, db Database) {
	// blueBg := color.New(color.BgBlue).SprintFunc()
	blueFg := color.New(color.FgBlue).SprintFunc()
	whiteFg := color.New(color.FgWhite).SprintFunc()

	fetchedPages := qr.Results
	maxTitleLength := 1
	for _, page := range fetchedPages {
		titleLength := len(page.Properties.Name.Title[0].Text.Content)
		if titleLength > maxTitleLength {
			maxTitleLength = titleLength
		}
	}

	// fmt.Printf("\n------------\n%s [%s]\n------------\n\n", blueFg(db.Title[0].Text.Content), db.Url)
	fmt.Printf("\n%s: %s\n", "Board ", db.Title[0].Text.Content)
	fmt.Printf("%s: %s\n", "Status", "Not started")
	fmt.Printf("%s: %s\n\n", "URL   ", db.Url)

	var flag uint = 0
	// flag := tabwriter.Debug
	writer := tabwriter.NewWriter(os.Stdout, 8, 8, 3, ' ', flag)
	defer writer.Flush()

	fmt.Fprintf(writer, "%s\t%s\t%s\n", whiteFg("Title"), "Created", "ID")
	fmt.Fprintf(writer, "%s\t%s\t%s\n", whiteFg("-----"), "-------", "--")
	for i := range fetchedPages {
		fmt.Fprintf(writer, "%s\t%s\t%s\n",
			blueFg(fetchedPages[i].Properties.Name.Title[0].Text.Content),
			fetchedPages[i].CreatedTime.Format("2006-01-02 15:04:05"),
			fetchedPages[i].Id,
		)
	}
}
