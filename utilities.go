package main

import "encoding/json"

func padding(s string) string {
	return " " + s + " "
}

// prettyPrint to print struct in a readable way
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
