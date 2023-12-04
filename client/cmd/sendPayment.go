/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sendPaymentCmd represents the sendPayment command
var sendPaymentCmd = &cobra.Command{
	Use:   "sendPayment",
	Short: "Send a payment to a node",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// set destinationNodePublicKey to the public key of the node you want to send a payment to
		destinationNodePublicKey := "bcrt1qna0pup6atlfxdspxhlxsvh4lt2a30qezcra43c"
		outgoingPayment, err := Client.SendPayment(NodeId, destinationNodePublicKey, 10000, 1000, 60)
		if err != nil {
			log.Printf("send payment failed: %v", err)
			return
		}
		log.Printf("Payment sent with payment id: %v\n", outgoingPayment.Id)
	},
}

func init() {
	rootCmd.AddCommand(sendPaymentCmd)
}
