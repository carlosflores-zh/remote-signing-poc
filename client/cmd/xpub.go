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
	Short: "Derive a public key used for node",
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
