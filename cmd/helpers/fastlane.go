package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Enviroment struct {
	Name  string
	Value string
}

func AppendEnvAtFastlane(envPath string, envs map[string]string) {
	if !FileExists(envPath) {
		fmt.Println("Create a .env fastlane before, or add manually the envs")
		return
	}

	var content = "#" + envs["APP_NAME"] + "\n"

	for key, value := range envs {
		content += key + "=" + value + "\n"
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

	fmt.Println("Envs adicionadas no final do arquivo .env, localizado em: " + Message(envPath, "yellow") + "\n")
}

func ConfirmEnvironmentVariables(environmentVariables map[string]string) {

	fmt.Println(Message("Before proceeding, please review the values of these environment variables carefully:\n", "blue"))
	for key, value := range environmentVariables {
		fmt.Println(key + "=" + Message(value, "yellow"))

	}
	fmt.Println(Message("\nAre the environment variables correctly? (Y/N):", "blue"))

	fmt.Print()
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var value = strings.TrimSuffix(line, "\n")
	if strings.EqualFold(value, "y") {
		return
	} else {
		os.Exit(0)
	}

}
