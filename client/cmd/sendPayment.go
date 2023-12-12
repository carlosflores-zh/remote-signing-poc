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
		if len(args) != 1 {
			log.Printf("Please provide a destination node public key")
			return
		}

		// set destinationNodePublicKey to the public key of the node you want to send a payment to
		destinationNodePublicKey := args[0]
		outgoingPayment, err := Client.SendPayment(NodeId, destinationNodePublicKey, 100, 1000, 10000)
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
