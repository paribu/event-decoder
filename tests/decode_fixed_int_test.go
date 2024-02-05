package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeFixedInt(t *testing.T) {
	abiFile := "decode_fixed_int_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x4e752b46070d13ba82917e2ec6cf2673d41ade08206f89e8f663e180191c492b"),
		},
		Data: "7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []ArrayParameter{
		{
			Type: "int256[2]",
			Name: "newValue",
			Value: []string{"57896044618658097711785492504343953926634992332820282019728792003956564819967",
				"57896044618658097711785492504343953926634992332820282019728792003956564819967"},
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
