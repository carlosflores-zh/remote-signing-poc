package cmd

import (
	"github.com/lightsparkdev/go-sdk/objects"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getTxsCmd represents the getTxs command
var getBlockchainBalance = &cobra.Command{
	Use:   "getBlockchainBalance",
	Short: "Prints the blockchain balance of the account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		networks := []objects.BitcoinNetwork{Network}
		blockchainBalance, err := Account.GetBlockchainBalance(Client.Requester, &networks, nil)
		if err != nil {
			log.Printf("get transactions failed: %v", err)
			return
		}

		log.Printf("You have blockchain balance: %v %v.\n", blockchainBalance.AvailableBalance.OriginalValue, blockchainBalance.AvailableBalance.OriginalUnit.StringValue())
	},
}

func init() {
	rootCmd.AddCommand(getBlockchainBalance)
}
