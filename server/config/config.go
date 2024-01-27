package config

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	HTTPPort int `json:"http_port"`
}

var singleton *configuration

func Load() error {
	bytesJson, err := os.ReadFile("../config/config.json")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo JSON: %s", err.Error())
	}

	singleton = &configuration{}
	err = json.Unmarshal(bytesJson, singleton)
	if err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err.Error())
	}
	return nil
}

func Get() *configuration {
	return singleton
	// return singleton
}
