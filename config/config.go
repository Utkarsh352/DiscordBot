package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	fmt.Printf("Token: %s\n", Token)		
	fmt.Printf("BotPrefix: %s\n", BotPrefix)
	fmt.Println("Config loaded")
	return nil
}
