package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createInvoiceCmd represents the createInvoice command
var createInvoiceCmd = &cobra.Command{
	Use:   "createInvoice",
	Short: "Create a test mode invoice",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		testInvoice, err := Client.CreateTestModeInvoice(NodeId, 250000, nil, nil)
		if err != nil {
			log.Printf("create test invoice failed: %v", err)
			return
		}

		log.Printf("Invoice created: %v\n", *testInvoice)
	},
}

func init() {
	rootCmd.AddCommand(createInvoiceCmd)
}
