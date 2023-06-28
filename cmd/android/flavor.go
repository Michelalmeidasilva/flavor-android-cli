package android

import (
	"android-cli/cmd/helpers"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Environments struct {
	category map[string]helpers.Enviroment
}

func AppendFlavorBuildGradle(buildGradlePath string, environmentVariables map[string]string) error {
	destinationContent, err := os.ReadFile(buildGradlePath)

	if err != nil {
		log.Fatal(err, ", please verify your flag values.")
	}

	productFlavor := fmt.Sprintf("\n\t\t%s  \t{\n\t\t\tdimension \"brand\"\n\t\t\tapplicationId \"%s\"\n\t\t\tresValue \"string\", \"build_config_package\", \"%s\"\n\t\t}\n", environmentVariables["APP_FLAVOR"], environmentVariables["BUNDLE_ID"], environmentVariables["PACKAGE_SRC"])

	var fileContent = string(destinationContent)

	productFlavorsIndex := strings.Index(fileContent, "productFlavors")
	if productFlavorsIndex == -1 {
		fmt.Println("Unable to find the productFlavors block in the file.")
	}

	modifiedContent := fileContent[:productFlavorsIndex+16] + productFlavor + fileContent[productFlavorsIndex+16:]

	err = ioutil.WriteFile(buildGradlePath, []byte(modifiedContent), 0644)

	if err != nil {
		return err
	}

	return nil

}

func CopyFlavorFolder(sourceFolder string, destinationFolder string, appFlavor string, deepLinking string) error {

	return helpers.CopyFolder(sourceFolder, destinationFolder, func(content []byte, fileName string) []byte {
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

func NewAndroidFlavor(environmentVariables map[string]string, pathToAndroidFolder string) {
	if pathToAndroidFolder != "" {
		var buildGradlePath = pathToAndroidFolder + "/app/build.gradle"
		var sourceFolderToCopy = pathToAndroidFolder + "/app/src/example"
		var destination = pathToAndroidFolder + "/app/src/" + environmentVariables["APP_FLAVOR"]
		var keystorePath = pathToAndroidFolder + "/app/keystores/"
		var appFlavor = environmentVariables["APP_FLAVOR"]

		//Implementar uma FSM e adicionar passo a passo, ver como fazer a leitura do usuário
		AppendFlavorBuildGradle(buildGradlePath, environmentVariables)
		CopyFlavorFolder(sourceFolderToCopy, destination, appFlavor, environmentVariables["DEEP_LINKING_TAG"])
		GenerateKeystore(keystorePath+environmentVariables["APP_FLAVOR"]+".keystore", environmentVariables["APP_KEY_ALIAS"], environmentVariables["APP_KEY_PASSWORD"], environmentVariables["APP_KEY_STORE_PASSWORD"])
		helpers.AppendEnvAtFastlane(pathToAndroidFolder+"/fastlane/.env", environmentVariables)
		helpers.CreateAndroidImages(environmentVariables["ICON_LAUNCHER_PATH"], destination+"/res/")

	} else {
		fmt.Printf("You need to provide a android path as first argument")
	}
}
