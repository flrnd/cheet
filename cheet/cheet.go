package cheet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Shortcut struct {
	Keys        string `json:"keys"`
	Description string `json:"description"`
}

type Sheet struct {
	Title     string `json:"title"`
	Shortcuts []Shortcut
}

func ParseFile() {
	homeDir, _ := os.UserHomeDir()
	fullPath := homeDir + "/.config/cheet/shortcuts.json"
	jsonFile, err := os.Open(fullPath)

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	var sheet Sheet

	json.Unmarshal(byteValue, &sheet)

	fmt.Println(sheet.Title)
	for _, s := range sheet.Shortcuts {
		fmt.Printf("%-12s : %s\n", s.Keys, s.Description)
	}
}
