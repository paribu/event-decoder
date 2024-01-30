package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeBool(t *testing.T) {
	abiFile := "decode_bool_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xd42752c19743102375aecef6d4643ae70cf8dff6d48cb2d31bbfba98142c5c24"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000001",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "bool",
			Name:  "newValue",
			Value: "true",
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
