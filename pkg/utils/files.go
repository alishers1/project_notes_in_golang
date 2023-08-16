package handlers

import (
	"encoding/json"
	"log"
	"notes/internal/handlers"
	"notes/pkg/models"
	"os"
)

var FileName = "notes.json"

func LoadNotesFromFile() {
	data, err := os.ReadFile(FileName)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(data, &models.Notes)
	if err != nil {
		log.Println(err)
		return
	}

	lastElement := models.Notes[len(models.Notes)-1]
	handlers.LastID = lastElement.ID
}

func SaveNotesToFile() {
	data, err := json.MarshalIndent(models.Notes, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	err = os.WriteFile(FileName, data, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}
