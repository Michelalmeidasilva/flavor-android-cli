package helper

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CopyFolder(sourceFolder string, destinationFolder string, executeFileOperation func(content []byte, fileName string) []byte) error {
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

	return nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}
