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
	Short: "Withdraw sats from a node, right now withdraws all AvailableToSendBalance from our remote signing node",
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
