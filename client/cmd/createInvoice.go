package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createInvoiceCmd
var createInvoiceCmd = &cobra.Command{
	Use:   "createInvoice",
	Short: "create Invoice for receiving payment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		memo := "payment"
		expiry := int32(300000)
		invoice, err := Client.CreateInvoice(NodeId, 10000, &memo, nil, &expiry)
		if err != nil {
			log.Printf("get node wallet failed: %v", err)
			return
		}

		log.Println("Invoice created: ", invoice.Data.EncodedPaymentRequest)
	},
}

func init() {
	rootCmd.AddCommand(createInvoiceCmd)
}
