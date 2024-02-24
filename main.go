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
	// fmt.Println(APIKEY)
	randomString := lib.GenerateRandomString()
	fmt.Println(randomString)
	encryptedKey := lib.Encrypt(APIKEY, randomString)
	fmt.Println(encryptedKey)
	decryptedKey := lib.Decrypt(encryptedKey, randomString)
	fmt.Println(decryptedKey)

}
