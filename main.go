package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
)

func setup() {
	var apiKey string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter your API-KEY").
				Prompt("🔑:").
				Value(&apiKey),
		))
	form.Run()

	processedString := fmt.Sprintf("APIKEY=%s", apiKey)
	env, _ := godotenv.Unmarshal(processedString)
	err := godotenv.Write(env, "./.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No config found! Lets set this up!")
		setup()
	}
}
