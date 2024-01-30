package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

type BoolArrayParameter struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value []bool `json:"value"`
}

func TestDecodeBoolArray(t *testing.T) {
	abiFile := "decode_bool_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xf7429222ad385bc52dc0cc1a37f9a54681dac04131f3c12aab17d388c66cd5d9"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000",
	}

	expectedParameters := []BoolArrayParameter{
		{
			Type:  "bool[]",
			Name:  "newValue",
			Value: []bool{true, false},
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

		var decodedBoolArray struct{ Values []bool }
		if err := json.Unmarshal([]byte(decoded.Value), &decodedBoolArray); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedBoolArray.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedBoolArray.Values)
		}
	}
}
