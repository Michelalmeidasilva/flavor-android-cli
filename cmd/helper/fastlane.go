package helper

import (
	"fmt"
	"os"
)

type Enviroment struct {
	Name  string
	Value string
}

func AppendEnvAtFastlane(envPath string, envs []Enviroment) {

	if !FileExists(envPath) {
		fmt.Println("Create a .env fastlane file and add manually the envs")
		return
	}

	var content = "#" + envs[0].Value + "\n"

	for _, value := range envs {
		content += value.Name + "=" + value.Value + "\n"
	}

	// Open the file in append mode, create if it doesn't exist
	file, err := os.OpenFile(envPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Append the content to the end of the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	fmt.Println("Envs adicionadas no final do arquivo .env" + envPath)
}