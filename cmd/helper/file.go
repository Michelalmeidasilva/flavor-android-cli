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

func AppendFlavorBuildGradle(newProductFlavors string, buildGradlePath string) error {

	destinationContent, err := os.ReadFile(buildGradlePath)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	var fileContent = string(destinationContent)

	productFlavorsIndex := strings.Index(fileContent, "productFlavors")
	if productFlavorsIndex == -1 {
		fmt.Println("Unable to find the productFlavors block in the file.")
	}

	modifiedContent := fileContent[:productFlavorsIndex+16] + newProductFlavors + fileContent[productFlavorsIndex+16:]

	err = ioutil.WriteFile(buildGradlePath, []byte(modifiedContent), 0644)

	if err != nil {
		return err
	}

	return nil

}

func CopyFlavorFolder(sourceFolder string, destinationFolder string, appFlavor string, deepLinking string) error {

	return CopyFolder(sourceFolder, destinationFolder, func(content []byte, fileName string) []byte {
		var newContent string = string(content)
		switch fileName {
		case "AndroidManifest.xml":
			newContent = strings.ReplaceAll(newContent, "exampleApp-9", deepLinking)

			break
		case "strings.xml":
			newContent = strings.ReplaceAll(newContent, "Example App", appFlavor)
			break

		default:
			break
		}

		return []byte(newContent)
	})
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

// func CopyFile(destinationPath string, sourceFile string) error {

// }
