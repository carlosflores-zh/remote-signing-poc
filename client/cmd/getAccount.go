/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getAccountCmd represents the getAccount command
var getAccountCmd = &cobra.Command{
	Use:   "getAccount",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		log.Printf("Your account name is: %v.\n", *Account.Name)

	},
}

func init() {
	rootCmd.AddCommand(getAccountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAccountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAccountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
