package helpers

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CopyFolder(sourceFolder string, destinationFolder string, executeFileOperation func(content []byte, fileName string) []byte) error {
	fmt.Println("\nArquivos criados em:")

	err := filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(sourceFolder, path)
		if err != nil {
			return err
		}

		destinationPath := filepath.Join(destinationFolder, relativePath)

		if info.IsDir() {
			err = os.MkdirAll(destinationPath, info.Mode())
			if err != nil {
				return err
			}

		} else {

			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			destinationFile, err := os.Create(destinationPath)
			if err != nil {
				return err
			}
			defer destinationFile.Close()

			_, err = io.Copy(destinationFile, sourceFile)

			if err != nil {
				return err
			}

			if executeFileOperation != nil {
				destinationContent, err := os.ReadFile(destinationPath)
				if err != nil {
					fmt.Println("Error", err)
					return err
				}

				var fileName = destinationFile.Name()[strings.LastIndex(destinationFile.Name(), "/")+1:]
				var modifiedContent = executeFileOperation(destinationContent, fileName)

				fmt.Println("\t" + Message(destinationPath, "yellow"))

				err = ioutil.WriteFile(destinationPath, []byte(modifiedContent), 0644)

				if err != nil {
					return err
				}

				return nil
			}

		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("")

	return nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func Message(message string, color string) string {
	var Yellow = "\033[33m"
	var Reset = "\033[0m"
	var Blue = "\033[34m"
	var Red = "\033[31m"

	if color == "yellow" {
		return (Yellow + message + Reset)
	}

	if color == "blue" {
		return (Blue + message + Reset)
	}

	if color == "red" {
		return (Red + message + Reset)
	}

	return message
}
