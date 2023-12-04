package cmd

import (
	"github.com/lightsparkdev/go-sdk/objects"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getNodesCmd represents the getNodes command
var getNodesCmd = &cobra.Command{
	Use:   "getNodes",
	Short: "Get all nodes for the account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		networks := []objects.BitcoinNetwork{Network}
		nodes, err := Account.GetNodes(Client.Requester, nil, &networks, nil, nil)
		if err != nil {
			log.Printf("get nodes failed: %v", err)
			return
		}

		for _, node := range nodes.Entities {
			balances := node.GetBalances()
			log.Printf("nodes: %s \n", node.GetId())
			log.Printf("Balance: %v %v \n", balances.AvailableToSendBalance.OriginalValue, balances.AvailableToSendBalance.OriginalUnit.StringValue())
			log.Printf("Balance: %v %v \n", balances.OwnedBalance.OriginalValue, balances.OwnedBalance.OriginalUnit.StringValue())
			log.Printf("Balance: %v %v \n \n", balances.AvailableToWithdrawBalance.OriginalValue, balances.AvailableToWithdrawBalance.OriginalUnit.StringValue())
		}

	},
}

func init() {
	rootCmd.AddCommand(getNodesCmd)
}
