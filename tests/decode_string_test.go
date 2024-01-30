package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeString(t *testing.T) {
	abiFile := "decode_string_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x402ddd519dec94a9c869d15434d5df9d0be41d3d937c9dbe2157489dc9a2be0d"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002a30786135303864643837356631306333336335326138616262323065313666633638653938316631383600000000000000000000000000000000000000000000",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "string",
			Name:  "newValue",
			Value: "0xa508dD875f10C33C52a8abb20E16fc68E981F186",
		},
	}

	parameters, err := decodeEventWithABI(abiFile, encodedEvent)
	if err != nil {
		t.Error(err)
	}

	for i, decoded := range parameters {
		expected := expectedParameters[i]

		if expected.Name != decoded.Name {
			t.Errorf("expected name %s, got %s", expected.Name, decoded.Name)
		}
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}

		decoded.Value = common.HexToAddress(decoded.Value).Hex()
		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}
