package android

import (
	"fmt"
	"log"
	"os/exec"
)

func GenerateKeystore(appFlavor string, keyAlias string, keyPassword string, keystorePassword string) {
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

	fmt.Printf("cmd: %v\n", cmd)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error to create keystore, verify your environment variables.")
	} else {
		fmt.Println("Keystore sucessfully created!")

	}

}