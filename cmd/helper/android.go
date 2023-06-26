package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
