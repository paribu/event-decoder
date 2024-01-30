package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeAddressArray(t *testing.T) {
	abiFile := "decode_address_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x0d98a1d69ce080b062fb35b76acc334c3ac51355ba51240273ee4e69fc1c4667"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000b4a925bae55743acf3dc65a8de0b9507f0491617000000000000000000000000288071244112050c93389a950d02c9e626d611dd",
	}

	expectedParameters := []ArrayParameter{
		{
			Type: "address[]",
			Name: "newValue",
			Value: []string{
				"0xB4a925BAe55743AcF3Dc65a8de0b9507F0491617",
				"0x288071244112050c93389A950d02c9E626D611dD"},
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

		var decodedArray Array
		if err := json.Unmarshal([]byte(decoded.Value), &decodedArray); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedArray.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedArray.Values)
		}
	}
}
