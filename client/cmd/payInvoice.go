/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// payInvoiceCmd represents the payInvoice command
var payInvoiceCmd = &cobra.Command{
	Use:   "payInvoice",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating a test mode invoice...")
		testInvoice, err := Client.CreateTestModeInvoice(NodeId, 250000, nil, nil)
		if err != nil {
			log.Printf("create test invoice failed: %v", err)
			return
		}

		log.Printf("Invoice created: %v\n", *testInvoice)

		outgoingPayment, err := Client.PayInvoice(NodeId, *testInvoice, 1000, 60, nil)
		if err != nil {
			log.Printf("pay invoice failed: %v", err)
			return
		}

		log.Printf("Invoice paid with payment id: %v\n", outgoingPayment.Id)
	},
}

func init() {
	rootCmd.AddCommand(payInvoiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// payInvoiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// payInvoiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
