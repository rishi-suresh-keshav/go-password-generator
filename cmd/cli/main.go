package main

import (
	"fmt"
	"os"

	gpg "github.com/rishi-suresh-keshav/go-password-generator/lib"
	"github.com/spf13/cobra"
)

var length int32

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates password with random characters",
	Long: `generates password with random characters
Optional flag --length will generate password with specified value`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt32("length")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pg := gpg.NewPasswordGenerator()
		if length > 0 {
			pg.WithLength(length)
		}

		password, err := pg.Generate()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Here is your random password of length %d: %s", pg.Length(), password))
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpg",
	Short: "A cli app to generate password containing random characters",
	Long:  `An app to generate password containing random characters`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func main() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().Int32VarP(&length, "length", "l", gpg.DefaultPasswordLength, "specify the required password length")
	rootCmd.AddCommand(generateCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
