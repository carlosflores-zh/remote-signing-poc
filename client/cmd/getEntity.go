/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/lightsparkdev/go-sdk/objects"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getEntityCmd represents the getEntity command
var getEntityCmd = &cobra.Command{
	Use:   "getEntity [entityID]",
	Short: "Get entity by ID, right now OutgoingPayment is the only supported entity type",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("getEntity requires an entity ID")
			return
		}
		entityID := args[0]

		entity, err := Client.GetEntity(entityID)
		if err != nil {
			fmt.Printf("get entity failed: %v", err)
		}

		log.Printf("Entity: %s > %s > %s \n", (*entity).GetTypename(), (*entity).GetId(), (*entity).(objects.OutgoingPayment).GetStatus().StringValue())
		log.Printf("Entity struct: %+v\n", *entity)
	},
}
