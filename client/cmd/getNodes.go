/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lightsparkdev/go-sdk/objects"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getNodesCmd represents the getNodes command
var getNodesCmd = &cobra.Command{
	Use:   "getNodes",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		networks := []objects.BitcoinNetwork{Network}
		nodeIDs := &[]string{NodeId}

		nodes, err := Account.GetNodes(Client.Requester, nil, &networks, nodeIDs, nil)
		if err != nil {
			log.Printf("get nodes failed: %v", err)
			return
		}

		for _, node := range nodes.Entities {
			balances := node.GetBalances()

			log.Printf("NodeID: %v Balance: %v %v \n", node.GetId(), balances.AvailableToSendBalance.OriginalValue, balances.AvailableToSendBalance.OriginalUnit.StringValue())
			log.Printf("NodeID: %v Balance: %v %v \n", node.GetId(), balances.OwnedBalance.OriginalValue, balances.OwnedBalance.OriginalUnit.StringValue())
			log.Printf("NodeID: %v Balance: %v %v \n", node.GetId(), balances.AvailableToWithdrawBalance.OriginalValue, balances.AvailableToWithdrawBalance.OriginalUnit.StringValue())

		}

	},
}

func init() {
	rootCmd.AddCommand(getNodesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNodesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNodesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
