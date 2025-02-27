package config

import (
	"encoding/json"
	"log"
	"os"
)

type Messages struct {
	InternalServerError string `json:"INTERNAL_SERVER_ERROR"`
	BadLocationName     string `json:"BAD_LOCATION_NAME"`
	InvalidDiscordCode  string `json:"INVALID_DISCORD_CODE"`
}

var msg Messages

func init() {
	file, err := os.Open("data/messages.json")
	if err != nil {
		log.Fatal("Cannot open messages.json : " + err.Error())
	}

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&msg); err != nil {
		log.Fatal("Cannot decode message.json : " + err.Error())
	}
}

func GetMessages() *Messages {
	return &msg
}
