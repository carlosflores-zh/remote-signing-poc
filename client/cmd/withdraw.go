/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lightsparkdev/go-sdk/objects"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// withdrawCmd represents the withdraw command
var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		networks := []objects.BitcoinNetwork{Network}
		nodes, err := Account.GetNodes(Client.Requester, nil, &networks, nil, nil)
		if err != nil {
			log.Printf("get nodes failed: %v", err)
			return
		}

		for _, node := range nodes.Entities {
			if node.GetId() == NodeId {
				// TODO receive bitcoin address from command line
				bitcoinAddress := "bcrt1qna0pup6atlfxdspxhlxsvh4lt2a30qezcra43c"
				// RequestWithdrawal receives sats
				log.Printf("balance to withdraw: %v", node.GetBalances().AvailableToSendBalance.OriginalValue/1000)
				withdrawalRequest, err := Client.RequestWithdrawal(NodeId, node.GetBalances().AvailableToSendBalance.OriginalValue/1000, bitcoinAddress, objects.WithdrawalModeWalletThenChannels)
				if err != nil {
					log.Printf("withdraw failed: %v", err)
					return
				}

				log.Printf("Withdrawal initiated with request id: %v\n", withdrawalRequest.Id)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(withdrawCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// withdrawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// withdrawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
