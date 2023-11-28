/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fundCmd represents the fundNode command
var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Execute fundNode operation, to fund a node with 10,000,000 sats in testnet",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		depositAmount, err := Client.FundNode(NodeId, 10000000) // 10,000,000 sats
		if err != nil {
			log.Printf("fund node failed: %v", err)
			return
		}

		log.Printf("Amount funded: %v %v \n", depositAmount.OriginalValue, depositAmount.OriginalUnit.StringValue())
	},
}
