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
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		depositAmount, err := Client.FundNode(NodeId, 10000000)
		if err != nil {
			log.Printf("fund node failed: %v", err)
			return
		}

		log.Printf("Amount funded: %v %v \n", depositAmount.OriginalValue, depositAmount.OriginalUnit.StringValue())
	},
}

func init() {
	rootCmd.AddCommand(fundCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fundCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fundCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
