package cmd

import (
	"os"
	"strings"

	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/lightsparkdev/go-sdk/services"
	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
	log "github.com/sirupsen/logrus"
)

var (
	NodeId  string
	Network objects.BitcoinNetwork
	Client  *services.LightsparkClient
	Seed    []byte
	Account *objects.Account
)

func Init() {
	var err error
	// MODIFY THOSE VARIABLES BEFORE RUNNING THE EXAMPLE
	apiClientID := os.Getenv("LS_CLIENT_ID")
	apiToken := os.Getenv("LS_TOKEN")
	baseUrl := os.Getenv("LS_BASE_URL")
	NodeId = os.Getenv("LS_NODE_ID")

	mnemonicSlice := strings.Split(os.Getenv("WORDS"), " ")
	Seed, err = lightspark_crypto.MnemonicToSeed(mnemonicSlice)
	if err != nil {
		log.Fatalf("mnemonic to seed failed: %v", err)
		return
	}

	Client = services.NewLightsparkClient(apiClientID, apiToken, &baseUrl)

	Account, err = Client.GetCurrentAccount()
	if err != nil {
		log.Fatalf("get current account failed: %v", err)
		return
	}

	Client.LoadNodeSigningKey(NodeId, *services.NewSigningKeyLoaderFromSignerMasterSeed(Seed, Network))

	log.Printf("Client: %+v\n", Client)
}
