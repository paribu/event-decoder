package tests

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeNestedStruct(t *testing.T) {
	abiFile := "decode_nested_struct_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x33d6a4306f041a3aabb18edf252ac6e69a54d0940111133bfeb0bc2ac03d583c"),
		},
		Data: "000000000000000000000000a508dd875f10c33c52a8abb20e16fc68e981f186",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "address",
			Name:  "newValue",
			Value: `{"A":"{\"A\":\"0xa508dD875f10C33C52a8abb20E16fc68E981F186\"}"}`,
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

		decoded.Type = strings.ReplaceAll(decoded.Type, "(", "")
		decoded.Type = strings.ReplaceAll(decoded.Type, ")", "")
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}

		if !reflect.DeepEqual(expected.Value, decoded.Value) {
			t.Errorf("expected value %v, got %v", expected.Value, decoded.Value)
		}
	}
}
