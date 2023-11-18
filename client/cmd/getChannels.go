/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getChannelsCmd represents the getChannels command
var getChannelsCmd = &cobra.Command{
	Use:   "getChannels",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		countc := int64(200)
		channels, err := Account.GetChannels(Client.Requester, Network, nil, nil, nil, &countc)
		if err != nil {
			log.Printf("get channels failed: %v", err)
			return
		}

		log.Printf("You have %v channels in total.\n", len(channels.Entities))

		for _, channel := range channels.Entities {
			log.Printf("ChannelID: %+v Status:%s LocalBalance:%d %s \n", channel.Id, channel.Status.StringValue(), (*(channel.LocalBalance)).OriginalValue, (*(channel.LocalBalance)).OriginalUnit.StringValue())
		}
	},
}

func init() {
	rootCmd.AddCommand(getChannelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getChannelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getChannelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
