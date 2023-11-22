// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package main

import (
	"fmt"
	"github.com/lightsparkdev/go-sdk/services"
	lightspark_crypto "github.com/lightsparkdev/lightspark-crypto-uniffi/lightspark-crypto-go"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/lightsparkdev/go-sdk/objects"
)

func main() {

	apiClientID := os.Getenv("LS_CLIENT_ID")
	apiToken := os.Getenv("LS_TOKEN")
	baseUrl := os.Getenv("LS_BASE_URL")
	nodeId := os.Getenv("LS_NODE_ID")

	mnemonicSlice := strings.Split(os.Getenv("WORDS"), " ")
	seed, err := lightspark_crypto.MnemonicToSeed(mnemonicSlice)
	if err != nil {
		log.Fatalf("mnemonic to seed failed: %v", err)
		return
	}

	// force to regtest
	network := objects.BitcoinNetworkRegtest
	client := services.NewLightsparkClient(apiClientID, apiToken, &baseUrl)

	account, err := client.GetCurrentAccount()
	if err != nil {
		log.Fatalf("get current account failed: %v", err)
		return
	}

	client.LoadNodeSigningKey(nodeId, *services.NewSigningKeyLoaderFromSignerMasterSeed(seed, network))

	entity, err := client.GetEntity("OutgoingPayment:018bf3e3-adfc-1d02-0000-3200942354bb")
	if err != nil {
		log.Fatalf("get entity failed: %v", err)
		return
	}

	log.Printf("entity: %+v", entity)

	countx := int64(10)
	transactionsConnection, err := account.GetTransactions(
		client.Requester,
		&countx,  // first
		nil,      // after
		nil,      // types
		nil,      // after_date
		nil,      // before_date
		&network, // bitcoin_network
		nil,      // lightning_node_id
		nil,      // statuses
		nil,      // exclude_failures
	)
	if err != nil {
		log.Fatalf("get payment requests failed: %v", err)
		return
	}

	for _, payment := range transactionsConnection.Entities {
		log.Printf("payment: %+v", payment)
	}

	fmt.Scanln()

	// Check your account's conductivity on REGTEST
	networks := []objects.BitcoinNetwork{objects.BitcoinNetworkRegtest}
	nodeIDs := &[]string{nodeId}

	nodes, err := account.GetNodes(client.Requester, nil, &networks, nodeIDs, nil)
	if err != nil {
		log.Printf("get nodes failed: %v", err)
		return
	}

	for _, node := range nodes.Entities {
		balances := node.GetBalances()

		log.Printf("NodeID: %v Balance: %v %v \n", node.GetId(), balances.AvailableToSendBalance.OriginalValue, balances.AvailableToSendBalance.OriginalUnit.StringValue())
		log.Printf("NodeID: %v Balance: %v %v \n", node.GetId(), balances.OwnedBalance.OriginalValue, balances.OwnedBalance.OriginalUnit.StringValue())
		log.Printf("NodeID: %v Balance: %v %v \n", node.GetId(), balances.AvailableToWithdrawBalance.OriginalValue, balances.AvailableToWithdrawBalance.OriginalUnit.StringValue())

	}

	bitcoinAddress := "bcrt1qna0pup6atlfxdspxhlxsvh4lt2a30qezcra43c"
	fmt.Printf("bitcoinAddress: %s", bitcoinAddress)
	// address, err := client.CreateNodeWalletAddress(nodeId)
	// if err != nil {
	// 	log.Printf("get node wallet failed: %v", err)
	// 	return
	// }
	//
	// bitcoinAddress = address

	// withdrawalRequest, err := client.RequestWithdrawal(nodeId, 90400, bitcoinAddress, objects.WithdrawalModeWalletThenChannels)
	// if err != nil {
	// 	log.Printf("withdraw failed: %v", err)
	// 	return
	// }
	//
	// log.Printf("Withdrawal initiated with request id: %v\n", withdrawalRequest.Id)
	//
	// fmt.Scanln()

	pointerNodeID := &nodeId
	var countc int64 = 200
	channelsConnection, err := account.GetChannels(client.Requester, network, pointerNodeID, nil, nil, &countc)
	if err != nil {
		log.Printf("get channels failed: %v", err)
		return
	}

	log.Printf("You have %v channels in total.\n", channelsConnection.Count)

	// depositAmount, err := client.FundNode(nodeId /* amountSats */, 10000000)
	// if err != nil {
	// 	log.Printf("fund node failed: %v", err)
	// 	return
	// }
	//
	// log.Printf("Amount funded: %v %v \n", depositAmount.OriginalValue, depositAmount.OriginalUnit.StringValue())
	// fmt.Scanln()

	channels, err := account.GetChannels(client.Requester, network, nil, nil, nil, &countc)
	if err != nil {
		log.Printf("get channels failed: %v", err)
		return
	}

	log.Printf("You have %v channels in total.\n", len(channels.Entities))

	for _, channel := range channels.Entities {
		log.Printf("ChannelID: %+v Status:%s LocalBalance:%d %s \n", channel.Id, channel.Status.StringValue(), (*(channel.LocalBalance)).OriginalValue, (*(channel.LocalBalance)).OriginalUnit.StringValue())
	}

	// ---------------------------------------------------------------------------------------------
	// Create an L1 address
	// log.Println("Creating an L1 address...")
	// address, err := client.CreateNodeWalletAddress(nodeId)
	// if err != nil {
	// 	log.Printf("get node wallet failed: %v", err)
	// 	return
	// }
	// log.Printf("Node wallet address created: %v\n", address)
	// log.Println()

	// // Check your nodes on REGTEST
	// nodesConnection, err := account.GetNodes(client.Requester, &count, &networks, nil, nil)
	// if err != nil {
	// 	log.Printf("get nodes failed: %v", err)
	// 	return
	// }
	// log.Printf("You have %v nodes in total.\n", nodesConnection.Count)
	//
	// for i, node := range nodesConnection.Entities {
	// 	log.Printf("#%v: %v with id %v\n", i, node.GetDisplayName(), node.GetId())
	// }

	var count int64 = 50
	transactionsConnection, err = account.GetTransactions(
		client.Requester,
		&count,   // first
		nil,      // after
		nil,      // types
		nil,      // after_date
		nil,      // before_date
		&network, // bitcoin_network
		nil,      // lightning_node_id
		nil,      // statuses
		nil,      // exclude_failures
	)
	if err != nil {
		log.Printf("get transactions failed: %v", err)
		return
	}
	log.Printf("You have %v transactions in total.\n", transactionsConnection.Count)

	var transactionId string
	for _, transaction := range transactionsConnection.Entities {
		transactionId = transaction.GetId()
		log.Printf(
			"    - %v at %v: %v %v (%v)\n",
			transactionId,
			transaction.GetCreatedAt(),
			transaction.GetAmount().OriginalValue,
			transaction.GetAmount().OriginalUnit.StringValue(),
			transaction.GetStatus().StringValue(),
		)
	}
	log.Println()

	// ----- invoice part

	// When testing paying invoice in test mode, a test invoice can be generated
	log.Println("Creating a test mode invoice...")
	testInvoice, err := client.CreateTestModeInvoice(nodeId, 2500000000, nil, nil)
	if err != nil {
		log.Printf("create test invoice failed: %v", err)
		return
	}
	encodedInvoice := *testInvoice
	log.Printf("Test invoice created: %v\n", *testInvoice)
	log.Println()

	// When testing paying invoice in test mode, a test invoice can be generated
	// log.Println("Creating an invoice...")
	// testInvoice, err := client.CreateInvoice(nodeId, 10000, nil, nil, nil)
	// if err != nil {
	// 	log.Printf("create test invoice failed: %v", err)
	// 	return
	// }
	//
	// log.Printf("Test invoice created: %v\n", *testInvoice)
	// fmt.Scanln()

	// Decode an encoded invoice
	log.Println("Decoding an encoded invoice...")
	decodedPaymentRequest, err := client.DecodePaymentRequest(encodedInvoice)
	if err != nil {
		log.Printf("decode invoice failed: %v", err)
		return
	}
	decodedInvoice, ok := (*decodedPaymentRequest).(objects.InvoiceData)
	if !ok {
		log.Printf("casting payment request to invoice failed")
		return
	}
	destinationNodePublicKey := *(decodedInvoice.Destination.GetPublicKey())
	log.Println("Decoded invoice...")
	log.Printf("    destination public key: %v\n", destinationNodePublicKey)
	log.Printf("    amount: %+v\n", decodedInvoice)
	log.Println("")
	fmt.Scanln()

	outgoingPayment, err := client.PayInvoice(nodeId, encodedInvoice, 1000, 60, nil)
	if err != nil {
		log.Printf("pay invoice failed: %v", err)
		return
	}
	log.Printf("Invoice paid with payment id: %v\n", outgoingPayment.Id)
	fmt.Scanln()

	// // ------ invoice part end

	//
	log.Println("Sending a payment...")
	outgoingPayment, err = client.SendPayment(nodeId, destinationNodePublicKey, 10000, 1000, 60)
	if err != nil {
		log.Printf("send payment failed: %v", err)
		return
	}
	log.Printf("Payment sent with payment id: %v\n", outgoingPayment.Id)
	fmt.Scanln()

	// Create an invoice
	// log.Println("Creating an invoice...")
	// invoice, err := client.CreateInvoice(nodeId, 100000, nil, nil, nil)
	// encodedInvoice = invoice.Data.EncodedPaymentRequest
	// if err != nil {
	// 	log.Printf("create invoice failed: %v", err)
	// 	return
	// }
	// log.Printf("Invoice created: %v\n", invoice.Data.EncodedPaymentRequest)
	// log.Println()

	// Get fee estimate for a node
	// log.Println("Getting fee estimate for a node...")
	// nodeFeeEstimate, err := client.GetLightningFeeEstimateForNode(nodeId, destinationNodePublicKey, 1000)
	// if err != nil {
	// 	log.Printf("getting fee estimate for node failed: %v", err)
	// 	return
	// }
	// log.Printf("Estimated fee for the node: %v %v\n", nodeFeeEstimate.FeeEstimate.OriginalValue, nodeFeeEstimate.FeeEstimate.OriginalUnit.StringValue())
	// log.Println()

	// Get fee estimate for an invoice
	// log.Println("Getting fee estimate for an invoice...")
	// invoiceFeeEstimate, err := client.GetLightningFeeEstimateForInvoice(nodeId, encodedInvoice, nil)
	// if err != nil {
	// 	log.Printf("getting fee estimate for invoice failed: %v", err)
	// 	return
	// }
	// log.Printf("Estimated fee for the invoice: %v %v\n", invoiceFeeEstimate.FeeEstimate.OriginalValue, invoiceFeeEstimate.FeeEstimate.OriginalUnit.StringValue())
	// If the node is in test mode, CreateTestModePayment can simulate a payment to the invoice

	newInvoice, err := client.CreateInvoice(nodeId, 10000, nil, nil, nil)
	if err != nil {
		log.Printf("create invoice failed: %v", err)
		return
	}

	testPayment, err := client.CreateTestModePayment(nodeId, newInvoice.Data.EncodedPaymentRequest, nil)
	if err != nil {
		log.Printf("simulating a test mode payment failed: %v", err)
		return
	}
	log.Printf("Invoice paid with a simulated payment %v\n", testPayment.Id)
}
