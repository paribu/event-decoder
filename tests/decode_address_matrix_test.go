package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeAddressMatrix(t *testing.T) {
	abiFile := "decode_address_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x82bca875692fb2e6d11ca902f16eb224600ba293462c00cbaf5cfba305de1ebd"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000020000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4000000000000000000000000ab8483f64d9c6d1ecf9b849ae677dd3315835cb200000000000000000000000000000000000000000000000000000000000000020000000000000000000000007e5f4552091a69125d5dfcb7b8c2659029395bdf0000000000000000000000007b23888aca8d32e6156c052eb9d3c28d8c346077",
	}

	expectedParameters := []MatrixParameter{{
		Type: "address[][]",
		Name: "newValue",
		Value: [][]string{
			{"0x5B38Da6a701c568545dCfcB03FcB875f56beddC4", "0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"},
			{"0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "0x7B23888ACA8d32e6156C052eb9d3c28D8C346077"},
		},
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

		var decodedMatrix Matrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}
