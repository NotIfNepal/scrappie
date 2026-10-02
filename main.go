package main

import (
	"log"
	"scrappie/src/helpers"
)

func main() {
	_, err := helpers.WgetFile("https://github.com/notifnepal/research/releases/latest/download/data.json", "downloads/data.json")
	if err != nil {
		log.Fatalf("failed to download the research data\n%s", err)
	}
	log.Print("Research file downloaded successfully!!")

	_, err = helpers.WgetFile("https://github.com/notifnepal/research/releases/latest/download/schema.json", "downloads/schema.json")
	if err != nil {
		log.Fatalf("failed to download the schema file for data processing\n%s", err)
	}
	log.Print("Schema file downloaded successfully!!")
}
