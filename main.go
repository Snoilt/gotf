package main

import (
	"fmt"
	"gotf/lib"

	"github.com/charmbracelet/huh"
)

func main() {
	var APIKEY string

	Textinput := huh.NewInput().
		Title("Enter your openAI API Key").
		Prompt("Key: ").
		Value(&APIKEY)

	Textinput.Run()
	fmt.Println(lib.GenerateRandomString(28))

}
