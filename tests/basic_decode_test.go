package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	event_decoder "github.com/paribu/event-decoder"
	"github.com/paribu/event-decoder/event"
)

var (
	_, b, _, _ = runtime.Caller(0)
	currentDir = filepath.Dir(b)
)

func TestBasicDecode(t *testing.T) {
	// Event we are going to send to decoder
	encodedEvent := &event.Event{
		Topics: []common.Hash{
			// Event Topic
			common.HexToHash("0x4b2e36342dfc5a22648f036210a3ef9d96202b7425615d0349559bae3192390f"),
			// First indexed parameter (name: account, type: address, value: 0xe2b5b0125c314e5967def2bde978d61b4114ef3f)
			common.HexToHash("0x000000000000000000000000e2b5b0125c314e5967def2bde978d61b4114ef3f"),
		},
		// Encoded: "Some data"
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000009536f6d6520646174610000000000000000000000000000000000000000000000",
	}

	// ABI for our event. We read it from a json file
	abiBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", currentDir, "basic_decode_abi.json"))
	if err != nil {
		t.Error(err)
	}

	// Parse read json.
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		t.Error(err)
	}

	// Expected results (parameters) after decoding
	expectedParameters := []event.Parameter{
		{
			Type:  "address",
			Name:  "account",
			Value: "0xe2B5B0125c314E5967DEf2bdE978d61b4114ef3F",
		},
		{
			Type:  "string",
			Name:  "data",
			Value: "Some data",
		},
	}

	// Actually decode event parameters
	parameters, err := event_decoder.Decode(encodedEvent, &contractABI)
	if err != nil {
		t.Error(err)
	}

	// Check results
	for i, decoded := range parameters {
		expected := expectedParameters[i]
		if expected.Name != decoded.Name {
			t.Errorf("expected name %s, got %s", expected.Name, decoded.Name)
		}
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}
		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}
