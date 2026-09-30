package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("wget", "--output-document=downloads/data.json", "https://github.com/notifnepal/research/releases/latest/download/data.json")

	_, err := cmd.Output()

	if err != nil {
		log.Fatalf("failed to download the research data\n%s", err)
	}
	log.Print("Research file downloaded successfully!!")
}
