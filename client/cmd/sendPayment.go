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
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		destinationNodePublicKey := ""
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendPaymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendPaymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
