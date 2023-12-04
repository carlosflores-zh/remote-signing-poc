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
	Short: "Prints the account name",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		log.Printf("Your account name is: %v.\n", *Account.Name)

	},
}

func init() {
	rootCmd.AddCommand(getAccountCmd)
}
