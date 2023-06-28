/*
Copyright © 2023 Michel Almeida da Silva
*/
package cmd

import (
	"android-cli/cmd/android"
	"os"

	"github.com/spf13/cobra"
)

var (
	BUNDLE_ID              string
	APP_FLAVOR             string
	BUILD_OUTPUT_TYPE      string
	APP_KEY_ALIAS          string
	APP_KEY_PASSWORD       string
	APP_KEY_STORE_PASSWORD string
	APP_NAME               string
	DEEP_LINKING_TAG       string
	PACKAGE_SRC            string
	ICON_LAUNCHER_PATH     string
)

var rootCmd = &cobra.Command{
	Use:     "android-cli",
	Version: "0.0.1",
	Short:   "This is a CLI for create a new ProductFlavor at Android project",
}

var androidFlavorCmd = &cobra.Command{
	Use: "create-android-flavor",
	Short: `This command create a new android flavor, you will need pass this following Flags:
	--BUNDLE_ID
	--APP_FLAVOR
	--BUILD_OUTPUT_TYPE
	--APP_KEY_ALIAS
	--APP_KEY_PASSWORD
	--APP_KEY_STORE_PASSWORD
	--APP_NAME
	--DEEP_LINKING_TAG`,
	Aliases: []string{"new-flavor"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		environmentVariables := map[string]string{
			"BUNDLE_ID":              BUNDLE_ID,
			"APP_FLAVOR":             APP_FLAVOR,
			"BUILD_OUTPUT_TYPE":      BUILD_OUTPUT_TYPE,
			"APP_KEY_ALIAS":          APP_KEY_ALIAS,
			"APP_KEY_PASSWORD":       APP_KEY_PASSWORD,
			"APP_KEY_STORE_PASSWORD": APP_KEY_STORE_PASSWORD,
			"APP_NAME":               APP_NAME,
			"DEEP_LINKING_TAG":       DEEP_LINKING_TAG,
			"PACKAGE_SRC":            PACKAGE_SRC,
			"ICON_LAUNCHER_PATH":     ICON_LAUNCHER_PATH,
		}

		android.NewAndroidFlavor(environmentVariables, args[0])
	},
	Example: `
1. Generate a new flavor app with icon and keystore:
- create-android-flavor ./examples/android --BUNDLE_ID="com.example.facebook" --APP_FLAVOR="facebook" --BUILD_OUTPUT_TYPE="AAB" --APP_KEY_ALIAS="my-key-alias" --APP_KEY_PASSWORD="my-password" --APP_KEY_STORE_PASSWORD="my-app-keystore-password" --APP_NAME="facebook" --DEEP_LINKING_TAG="facebookApp-8574" --PACKAGE_SRC="com.example" --ICON_LAUNCHER_PATH="examples/icone.png"`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

func init() {
	androidFlavorCmd.Flags().StringVar(&APP_FLAVOR, "APP_FLAVOR", "", "Android flavors allow you to create different variants of your application, typically used for different environments (e.g., development, staging, production). The --APP_FLAVOR flag specifies the flavor of the application to build.")
	androidFlavorCmd.Flags().StringVar(&BUNDLE_ID, "BUNDLE_ID", "", "The bundle ID (also known as package name) uniquely identifies an Android application. It is typically in reverse domain name format (e.g., com.example.myapp)")
	androidFlavorCmd.Flags().StringVar(&BUILD_OUTPUT_TYPE, "BUILD_OUTPUT_TYPE", "APK", "This flag determines the type of build output for the Android application. it can be APK or AAB")
	androidFlavorCmd.Flags().StringVar(&APP_KEY_ALIAS, "APP_KEY_ALIAS", "my-key-alias", "The key alias refers to the alias of the signing key used to sign the Android application. It is used during the app signing process.")
	androidFlavorCmd.Flags().StringVar(&APP_KEY_PASSWORD, "APP_KEY_PASSWORD", "", "The password associated with the signing key alias. This password is required to access and use the signing key during the signing process.")
	androidFlavorCmd.Flags().StringVar(&APP_KEY_STORE_PASSWORD, "APP_KEY_STORE_PASSWORD", "", "The password for the key store file that contains the signing key. The key store file is used to securely store the signing key and protect it from unauthorized access.")
	androidFlavorCmd.Flags().StringVar(&APP_NAME, "APP_NAME", "", "The name of the Android application. It is a user-friendly name that is displayed to users when they interact with the app.")
	androidFlavorCmd.Flags().StringVar(&DEEP_LINKING_TAG, "DEEP_LINKING_TAG", "", "Deep linking allows users to navigate directly to specific screens or content within an app. The --DEEP_LINKING_TAG flag specifies a tag or identifier associated with a deep link, which can be used to handle deep link URLs within the Android application.")
	androidFlavorCmd.Flags().StringVar(&PACKAGE_SRC, "PACKAGE_SRC", "", "")
	androidFlavorCmd.Flags().StringVar(&ICON_LAUNCHER_PATH, "ICON_LAUNCHER_PATH", "", "Path to image.png to generate a image")

	androidFlavorCmd.MarkFlagRequired("APP_FLAVOR")
	androidFlavorCmd.MarkFlagRequired("BUNDLE_ID")
	androidFlavorCmd.MarkFlagRequired("BUILD_OUTPUT_TYPE")
	androidFlavorCmd.MarkFlagRequired("APP_KEY_ALIAS")
	androidFlavorCmd.MarkFlagRequired("APP_KEY_PASSWORD")
	androidFlavorCmd.MarkFlagRequired("APP_KEY_STORE_PASSWORD")
	androidFlavorCmd.MarkFlagRequired("APP_NAME")
	androidFlavorCmd.MarkFlagRequired("DEEP_LINKING_TAG")
	androidFlavorCmd.MarkFlagRequired("PACKAGE_SRC")

	rootCmd.AddCommand(androidFlavorCmd)

}
