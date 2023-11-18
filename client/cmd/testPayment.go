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
	Short: "A brief description of your command",
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

func init() {
	rootCmd.AddCommand(testPaymentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testPaymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testPaymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
