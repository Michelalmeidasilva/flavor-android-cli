/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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

var androidFlavorCmd = &cobra.Command{
	Use:     "create-android-flavor",
	Short:   "This command create a new android flavor, you will need pass this following information:",
	Aliases: []string{"new-flavor"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// res := helper.(args[0])

		// helper.CopyFolder()
		// fmt.Println(res)
	},
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.AddCommand(reverseCmd)
	rootCmd.AddCommand(androidFlavorCmd)

}
