package android

import (
	"android-cli/cmd/helpers"
	"fmt"
	"log"
	"os/exec"
)

func GenerateKeystore(appFlavor string, keyAlias string, keyPassword string, keystorePassword string) {

	if helpers.FileExists(appFlavor) {
		log.Fatalln(helpers.Message("Error generating keystore file, this path '"+appFlavor+"' already exists a file with this name.", "red"))
		// os.Exit(1)
	}
	dname := "CN=, OU=, O=, L=, ST=, C="

	cmd := exec.Command("keytool",
		"-genkey",
		"-v",
		"-keystore", appFlavor,
		"-alias", keyAlias,
		"-keyalg", "RSA",
		"-keysize", "2048",
		"-validity", "10000",
		"-keypass", keyPassword,
		"-storepass", keystorePassword,
		"-dname", dname,
	)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error to create keystore, verify your environment variables.")
	} else {
		fmt.Println("Keystore sucessfully created at following path:" + helpers.Message(appFlavor, "yellow") + "\n")

	}

}
