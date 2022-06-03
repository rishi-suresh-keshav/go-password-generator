/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

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
			os.Exit(0)
		}

		pg := gpg.NewPasswordGenerator()
		if length > 0 {
			pg.WithLength(length)
		}

		password, err := pg.Generate()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Println(fmt.Sprintf("Here is your random password of length %d: %s", pg.Length(), password))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().Int32VarP(&length, "length", "l", gpg.DefaultPasswordLength, "specify the required password length")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
