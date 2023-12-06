package cmd

import (
	"github.com/lightsparkdev/go-sdk/objects"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

// withdrawCmd represents the withdrawal command
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

		log.Printf("You have %v nodes in total.\n", len(nodes.Entities))

		for _, node := range nodes.Entities {
			// splitting because ids don't match completely, just the last part (UUID)
			if strings.SplitAfter(node.GetId(), ":")[1] == strings.SplitAfter(NodeId, ":")[1] {
				// my prod personal one
				bitcoinAddress := "bc1qxz9udg6lvd8pvfezkqtvze5ppdj9tkkhszppgu"

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
}
