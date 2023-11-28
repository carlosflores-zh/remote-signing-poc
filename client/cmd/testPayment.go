/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// testPaymentCmd represents the testPayment command
var testPaymentCmd = &cobra.Command{
	Use:   "testPayment",
	Short: "Create an invoice and pay it with a simulated payment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		newInvoice, err := Client.CreateInvoice(NodeId, 10000, nil, nil, nil)
		if err != nil {
			log.Printf("create invoice failed: %v", err)
			return
		}

		log.Printf("Invoice created: %v\n", newInvoice.Data.EncodedPaymentRequest)

		testPayment, err := Client.CreateTestModePayment(NodeId, newInvoice.Data.EncodedPaymentRequest, nil)
		if err != nil {
			log.Printf("simulating a test mode payment failed: %v", err)
			return
		}

		log.Printf("Invoice paid with a simulated payment %v\n", testPayment.Id)
	},
}
