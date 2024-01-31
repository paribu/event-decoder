package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeIndexedAddress(t *testing.T) {
	abiFile := "decode_indexed_address_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x6bba86f485affce7736750bc725be237058c460106794e2c3e4a8fdd7283be33"),
		},
		Data: "000000000000000000000000a508dd875f10c33c52a8abb20e16fc68e981f186000000000000000000000000a508dd875f10c33c52a8abb20e16fc68e981f186",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "address[2]",
			Name:  "newValue",
			Value: []string{"0xa508dD875f10C33C52a8abb20E16fc68E981F186", "0xa508dD875f10C33C52a8abb20E16fc68E981F186"},
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
