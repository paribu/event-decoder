package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeFixedString(t *testing.T) {
	abiFile := "decode_fixed_string_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xe36dc9bef6a334cb29006d59c57601ebb866681e8383f43d883d1b621c0b68d1"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000002a30786135303864643837356631306333336335326138616262323065313666633638653938316631383600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002a30786135303864643837356631306333336335326138616262323065313666633638653938316631383600000000000000000000000000000000000000000000",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "string[2]",
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

		for j := 0; j < len(decodedArray.Values); j++ {
			if !reflect.DeepEqual(common.FromHex(expected.Value[j]), common.FromHex(decodedArray.Values[j])) {
				t.Errorf("expected value %v, got %v", expected.Value[j], decodedArray.Values[j])
			}
		}
	}
}
