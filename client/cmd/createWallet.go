/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		address, err := Client.CreateNodeWalletAddress(NodeId)
		if err != nil {
			log.Printf("get node wallet failed: %v", err)
			return
		}
		log.Printf("Node wallet address created: %v\n", address)
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
