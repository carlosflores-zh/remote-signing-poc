// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
)

type Config struct {
	ApiEndpoint     *string
	ApiClientId     string
	ApiClientSecret string
	WebhookSecret   string
	MasterSeed      []byte
}

func NewConfigFromEnv() (*Config, error) {
	mnemonicSlice := strings.Split(os.Getenv("WORDS"), " ")
	masterSeed, err := lightspark_crypto.MnemonicToSeed(mnemonicSlice)
	if err != nil {
		log.Fatalf("Invalid mnemonic: %s", err)
	}

	apiClientId := os.Getenv("LS_CLIENT_ID")
	apiClientSecret := os.Getenv("LS_TOKEN")
	webhookSecret := os.Getenv("LS_WEBHOOK_SECRET")
	apiEndpointStr := os.Getenv("LS_BASE_URL")
	
	log.Print("Loaded configuration:")
	log.Printf("  - API_CLIENT_ID: %s", showEmpty(apiClientId))
	log.Printf("  - API_CLIENT_SECRET: %s", showEmpty(fmt.Sprint(len(apiClientSecret))))
	log.Printf("  - WEBHOOK_SECRET: %s", showEmpty(fmt.Sprint(len(webhookSecret))))
	log.Printf("  - MASTER_SEED: %s", showEmpty(fmt.Sprint(len(masterSeed))))
	log.Printf("  - API_ENDPOINT: %s", showEmpty(apiEndpointStr))

	return &Config{
		ApiEndpoint:     &apiEndpointStr,
		ApiClientId:     apiClientId,
		ApiClientSecret: apiClientSecret,
		WebhookSecret:   webhookSecret,
		MasterSeed:      masterSeed,
	}, nil
}

func showEmpty(str string) string {
	if str == "" {
		return "<empty>"
	}

	return str
}
