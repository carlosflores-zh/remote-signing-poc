// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package remotesigning

import (
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcd/txscript"
	"github.com/lightsparkdev/go-sdk/webhooks"
	log "github.com/sirupsen/logrus"
	"regexp"
)

// Validator an interface which decides whether to sign or reject a remote signing webhook event.
type Validator interface {
	ShouldSign(webhookEvent webhooks.WebhookEvent) bool
}

type PositiveValidator struct{}

func (v PositiveValidator) ShouldSign(event webhooks.WebhookEvent) bool {
	subEventTypeStr := (*event.Data)["sub_event_type"].(string)

	// validate sub event type to reject if needed
	if subEventTypeStr == "DERIVE_KEY_AND_SIGN" {
		return true
	}

	return true
}

func GetPaymentHashFromScript(scriptHex string) (*string, error) {
	pattern := `OP_HASH160 ([a-fA-F0-9]{40}) OP_EQUALVERIFY`

	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		return nil, err
	}

	log.Printf("script: %s", script)

	disassembled, err := txscript.DisasmString(script)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(disassembled)
	if len(match) > 0 {
		return &match[1], nil
	} else {
		return nil, errors.New("No match found")
	}
}
