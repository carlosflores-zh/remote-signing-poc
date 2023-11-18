/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// xpubCmd represents the xpub command
var xpubCmd = &cobra.Command{
	Use:   "xpub",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		publicKey, err := lightspark_crypto.DerivePublicKey(Seed, 3, "m")
		if err != nil {
			log.Printf("get public key failed: %v", err)
			return
		}

		log.Printf("xpub: %s", publicKey)
	},
}

func init() {
	rootCmd.AddCommand(xpubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xpubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xpubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
