package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createWalletCmd
var getLnFeesCmd = &cobra.Command{
	Use:   "getLnFees",
	Short: "get LN Fees",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Printf("Please provide an invoice")
			return
		}

		amount := int64(0)
		lnFees, err := Client.GetLightningFeeEstimateForInvoice(NodeId, args[0], &amount)
		if err != nil {
			log.Printf("get node wallet failed: %v", err)
			return
		}

		log.Printf("Lightning Fees: %+v\n", lnFees)
	},
}

func init() {
	rootCmd.AddCommand(getLnFeesCmd)
}
