/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "android-cli",
	Version: "0.0.1",
	Short:   "This is a CLI for create a new ProductFlavor at Android project",
	Long: `
		A longer description that spans multiple lines and likely contains
		examples and usage of using your application. For example:
		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.
	`,
}

var (
	BUNDLE_ID              string
	APP_FLAVOR             string
	BUILD_OUTPUT_TYPE      string
	APP_KEY_ALIAS          string
	APP_KEY_PASSWORD       string
	APP_KEY_STORE_PASSWORD string
	APP_NAME               string
	DEEP_LINKING_TAG       string
)

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

		/**
		Validate a correct android folder ( maybe build.gradle or others files to validate if its correct)
		*/
		if args[0] != "" {
			fmt.Println("path to android folder", args[0])
		}
		fmt.Println(APP_FLAVOR)
		fmt.Println(BUNDLE_ID)
		fmt.Println(BUILD_OUTPUT_TYPE)
		fmt.Println(APP_KEY_ALIAS)
		fmt.Println(APP_KEY_PASSWORD)
		fmt.Println(APP_KEY_STORE_PASSWORD)
		fmt.Println(APP_NAME)
		fmt.Println(DEEP_LINKING_TAG)

	},
	Example: `
1. Generate a new flavor app:
- create-android-flavor path/to/android-folder  --BUNDLE_ID="com.example.facebook" --APP_FLAVOR="facebook" --BUILD_OUTPUT_TYPE="AAB" --APP_KEY_ALIAS="my-key-alias" --APP_KEY_PASSWORD="my-password" --APP_KEY_STORE_PASSWORD="my-app-keystore-password" --APP_NAME="facebook" --DEEP_LINKING_TAG="facebookApp-8574"`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.android-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	androidFlavorCmd.Flags().StringVar(&APP_FLAVOR, "APP_FLAVOR", "", "Android flavors allow you to create different variants of your application, typically used for different environments (e.g., development, staging, production). The --APP_FLAVOR flag specifies the flavor of the application to build.")
	androidFlavorCmd.Flags().StringVar(&BUNDLE_ID, "BUNDLE_ID", "", "The bundle ID (also known as package name) uniquely identifies an Android application. It is typically in reverse domain name format (e.g., com.example.myapp)")
	androidFlavorCmd.Flags().StringVar(&BUILD_OUTPUT_TYPE, "BUILD_OUTPUT_TYPE", "APK", "This flag determines the type of build output for the Android application. it can be APK or AAB")
	androidFlavorCmd.Flags().StringVar(&APP_KEY_ALIAS, "APP_KEY_ALIAS", "my-key-alias", "The key alias refers to the alias of the signing key used to sign the Android application. It is used during the app signing process.")
	androidFlavorCmd.Flags().StringVar(&APP_KEY_PASSWORD, "APP_KEY_PASSWORD", "", "The password associated with the signing key alias. This password is required to access and use the signing key during the signing process.")
	androidFlavorCmd.Flags().StringVar(&APP_KEY_STORE_PASSWORD, "APP_KEY_STORE_PASSWORD", "", "The password for the key store file that contains the signing key. The key store file is used to securely store the signing key and protect it from unauthorized access.")
	androidFlavorCmd.Flags().StringVar(&APP_NAME, "APP_NAME", "", "The name of the Android application. It is a user-friendly name that is displayed to users when they interact with the app.")
	androidFlavorCmd.Flags().StringVar(&DEEP_LINKING_TAG, "DEEP_LINKING_TAG", "", "Deep linking allows users to navigate directly to specific screens or content within an app. The --DEEP_LINKING_TAG flag specifies a tag or identifier associated with a deep link, which can be used to handle deep link URLs within the Android application.")

	androidFlavorCmd.MarkFlagRequired("APP_FLAVOR")
	androidFlavorCmd.MarkFlagRequired("BUNDLE_ID")
	androidFlavorCmd.MarkFlagRequired("BUILD_OUTPUT_TYPE")
	androidFlavorCmd.MarkFlagRequired("APP_KEY_ALIAS")
	androidFlavorCmd.MarkFlagRequired("APP_KEY_PASSWORD")
	androidFlavorCmd.MarkFlagRequired("APP_KEY_STORE_PASSWORD")
	androidFlavorCmd.MarkFlagRequired("APP_NAME")
	androidFlavorCmd.MarkFlagRequired("DEEP_LINKING_TAG")

	rootCmd.AddCommand(androidFlavorCmd)

}
