package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeAddressEvent(t *testing.T) {
	abiFile := "decode_address_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x8eef533f864799eb1d12894c25d39600718719b601ecc756b1311264486c385a"),
		},
		Data: "000000000000000000000000a508dd875f10c33c52a8abb20e16fc68e981f186",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "address",
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
		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}
